import {
  InventoryResponse,
  TransactionsResponse,
  AuthResponse,
  LoginRequest,
  RegisterRequest,
  User,
  Case,
  CaseOpenResult,
} from './types';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

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
    return authenticatedFetch('/user/profile');
  },

  updateProfile: async (data: Partial<User>): Promise<User> => {
    return authenticatedFetch('/user/profile', {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  },
};

export const inventoryAPI = {
  getInventory: async (showSold = false): Promise<InventoryResponse> => {
    return authenticatedFetch(`/inventory?show_sold=${showSold}`);
  },

  sellItem: async (itemId: string) => {
    return authenticatedFetch(`/inventory/${itemId}/sell`, {
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
    return publicFetch('/cases');
  },

  // this will get a specific case by id
  getCaseById: async (caseId: string): Promise<Case> => {
    // this is a public endpoint
    return publicFetch(`/cases/${caseId}`);
  },

  // this will open a case and return the result
  openCase: async (caseId: string): Promise<CaseOpenResult> => {
    return authenticatedFetch(`/cases/${caseId}/open`, {
      method: 'POST',
    });
  },
};
