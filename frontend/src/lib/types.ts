// User types
export interface User {
  id: string;
  email: string;
  username: string;
  casebucks: number;
  created_at: string;
}

// Skin types
export interface Skin {
  id: string;
  name: string;
  weapon_type: string;
  rarity: string;
  image_url: string;
  min_value: number;
  max_value: number;
}

// Inventory item
export interface InventoryItem {
  id: string;
  skin_id: string;
  float: number;
  condition: string;
  value: number;
  acquired_from: string;
  acquired_at: string;
  is_sold: boolean;
  sold_at: string | null;
  skin: Skin;
}

// Inventory response from API
export interface InventoryResponse {
  items: InventoryItem[];
  total_value: number;
  item_count: number;
  stats: Record<string, number>;
}

// Transaction types
export interface Transaction {
  id: string;
  type: string;
  amount: number;
  balance_before: number;
  balance_after: number;
  description: string;
  created_at: string;
  reference_id?: string;
}

export interface TransactionsResponse {
  transactions: Transaction[];
  count: number;
}

// Auth types
export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  email: string;
  username: string;
  password: string;
}

export interface AuthResponse {
  token: string;
  user: User;
  daily_reward_claimed?: boolean;
  daily_reward_amount?: number;
}

// case types
export interface Case {
  id: string;
  name: string;
  description: string;
  image_url: string;
  price: number;
  created_at: string;
  updated_at: string;
}

export interface CaseSkinOption extends Skin {
  drop_chance: number;
  drop_percentage: number;
}

export interface CaseDetails extends Case {
  skins: CaseSkinOption[];
}

// case opening result
export interface CaseOpenResult {
  message: string;
  case: Case;
  skin: Skin;
  float: number;
  condition: string;
  value: number;
  new_balance: number;
  inventory_id: string;
  transaction_id: string;
}

export interface PurchasedCase {
  id: string;
  user_id: string;
  case_id: string;
  is_opened: boolean;
  opened_at?: string;
  created_at: string;
  updated_at: string;
  case: Case;
}

export interface PurchasedCasesResponse {
  cases: PurchasedCase[];
  count: number;
}

export interface CaseBuyResult {
  message: string;
  purchased_case: PurchasedCase;
  case: Case;
  new_balance: number;
  transaction_id: string;
}

export interface PriceCheckResponse {
  provider: string;
  skin_name: string;
  suggested_usd: number;
  message: string;
}

// Rarity colors for UI
export const RARITY_COLORS: Record<string, string> = {
  'Consumer Grade': 'from-gray-400 to-gray-500',
  'Industrial Grade': 'from-blue-400 to-blue-500',
  'Mil-Spec': 'from-blue-500 to-blue-600',
  Restricted: 'from-purple-500 to-purple-600',
  Classified: 'from-pink-500 to-pink-600',
  Covert: 'from-red-500 to-red-600',
  'Rare Special': 'from-yellow-400 to-yellow-500',
};

export const RARITY_TEXT_COLORS: Record<string, string> = {
  'Consumer Grade': 'text-gray-400',
  'Industrial Grade': 'text-blue-400',
  'Mil-Spec': 'text-blue-500',
  Restricted: 'text-purple-500',
  Classified: 'text-pink-500',
  Covert: 'text-red-500',
  'Rare Special': 'text-yellow-400',
};

export const CONDITION_COLORS: Record<string, string> = {
  'Factory New': 'text-green-400',
  'Minimal Wear': 'text-blue-400',
  'Field-Tested': 'text-yellow-400',
  'Well-Worn': 'text-orange-400',
  'Battle-Scarred': 'text-red-400',
};
