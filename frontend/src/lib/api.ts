import {
  InventoryResponse,
  TransactionsResponse,
  AuthResponse,
  LoginRequest,
  RegisterRequest,
  User,
  Case,
  CaseOpenResult,
  CaseDetails,
  PurchasedCasesResponse,
  CaseBuyResult,
  PriceCheckResponse,
} from './types';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

const getConditionFromFloat = (float: number): string => {
  if (float < 0.07) return 'Factory New';
  if (float < 0.15) return 'Minimal Wear';
  if (float < 0.38) return 'Field-Tested';
  if (float < 0.45) return 'Well-Worn';
  return 'Battle-Scarred';
};

type RawInventoryItem = {
  float: number;
  created_at?: string;
  acquired_at?: string;
  condition?: string;
  sold_at?: string | null;
  [key: string]: unknown;
};

const getAuthToken = (): string | null => {
  if (typeof window === 'undefined') return null;
  return localStorage.getItem('auth_token');
};

export const setAuthToken = (token: string): void => {
  if (typeof window !== 'undefined') {
    localStorage.setItem('auth_token', token);
  }
};

export const removeAuthToken = (): void => {
  if (typeof window !== 'undefined') {
    localStorage.removeItem('auth_token');
  }
};

const authenticatedFetch = async (
  endpoint: string,
  options: RequestInit = {}
) => {
  const token = getAuthToken();

  if (!token) {
    throw new Error('Not authenticated. Please login.');
  }

  const response = await fetch(`${API_BASE_URL}${endpoint}`, {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`,
      ...options.headers,
    },
  });

  if (!response.ok) {
    const error = await response
      .json()
      .catch(() => ({ error: 'Request failed' }));
    throw new Error(error.error || `HTTP ${response.status}`);
  }

  return response.json();
};

const publicFetch = async (endpoint: string, options: RequestInit = {}) => {
  const response = await fetch(`${API_BASE_URL}${endpoint}`, {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options.headers,
    },
  });

  if (!response.ok) {
    const error = await response
      .json()
      .catch(() => ({ error: 'Request failed' }));
    throw new Error(error.error || `HTTP ${response.status}`);
  }

  return response.json();
};

export const authAPI = {
  login: async (credentials: LoginRequest): Promise<AuthResponse> => {
    const response = await publicFetch('/auth/login', {
      method: 'POST',
      body: JSON.stringify(credentials),
    });
    setAuthToken(response.token);
    return response;
  },

  register: async (data: RegisterRequest): Promise<AuthResponse> => {
    const response = await publicFetch('/auth/register', {
      method: 'POST',
      body: JSON.stringify(data),
    });
    setAuthToken(response.token);
    return response;
  },

  logout: () => {
    removeAuthToken();
  },

  isAuthenticated: (): boolean => {
    return !!getAuthToken();
  },
};

export const userAPI = {
  getProfile: async (): Promise<User> => {
    const response = await authenticatedFetch('/user/profile');
    return response.user;
  },

  updateProfile: async (data: Partial<User>): Promise<User> => {
    const response = await authenticatedFetch('/user/profile', {
      method: 'PUT',
      body: JSON.stringify(data),
    });
    return response.user;
  },
};

export const inventoryAPI = {
  getInventory: async (showSold = false): Promise<InventoryResponse> => {
    const response = await authenticatedFetch(`/inventory?show_sold=${showSold}`);
    const rawItems = (response.items || []) as RawInventoryItem[];
    const items = rawItems.map((item) => ({
      ...item,
      acquired_at: item.acquired_at || item.created_at,
      condition: item.condition || getConditionFromFloat(item.float),
      sold_at: item.sold_at || null,
    }));

    return {
      ...response,
      items,
    };
  },

  sellItem: async (itemId: string) => {
    return authenticatedFetch(`/inventory/${itemId}/sell`, {
      method: 'POST',
    });
  },

  getPurchasedCases: async (): Promise<PurchasedCasesResponse> => {
    return authenticatedFetch('/inventory/cases');
  },

  openPurchasedCase: async (purchasedCaseId: string): Promise<CaseOpenResult> => {
    return authenticatedFetch(`/inventory/cases/${purchasedCaseId}/open`, {
      method: 'POST',
    });
  },
};

export const transactionsAPI = {
  getTransactions: async (
    type?: string,
    limit = 50
  ): Promise<TransactionsResponse> => {
    const params = new URLSearchParams();
    if (type) params.append('type', type);
    params.append('limit', limit.toString());

    return authenticatedFetch(`/transactions?${params}`);
  },
};

// This is a API section for the case opening feature
export const casesAPI = {
  //this gets all available cases
  getAllCases: async (): Promise<Case[]> => {
    // public endpoint
    const response = await publicFetch('/cases');
    return response.cases || [];
  },

  // this will get a specific case by id
  getCaseById: async (caseId: string): Promise<CaseDetails> => {
    // this is a public endpoint
    return publicFetch(`/cases/${caseId}`);
  },

  // this will open a case and return the result
  openCase: async (caseId: string): Promise<CaseOpenResult> => {
    return authenticatedFetch(`/cases/${caseId}/open`, {
      method: 'POST',
    });
  },

  buyCase: async (caseId: string): Promise<CaseBuyResult> => {
    return authenticatedFetch(`/cases/${caseId}/buy`, {
      method: 'POST',
    });
  },
};

export const aiAPI = {
  priceCheck: async (skinName: string): Promise<PriceCheckResponse> => {
    return authenticatedFetch('/ai/price-check', {
      method: 'POST',
      body: JSON.stringify({ skin_name: skinName }),
    });
  },
};
