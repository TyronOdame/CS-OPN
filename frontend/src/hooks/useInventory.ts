'use client';

import { useState, useEffect } from 'react';
import { inventoryAPI } from '@/lib/api';
import { InventoryResponse } from '@/lib/types';

export const useInventory = (showSold = false) => {
  const [data, setData] = useState<InventoryResponse | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchInventory = async () => {
    try {
      setLoading(true);
      setError(null);
      const response = await inventoryAPI.getInventory(showSold);
      setData(response);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to load inventory');
      console.error('Inventory fetch error:', err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchInventory();
  }, [showSold]);

  const sellItem = async (itemId: string) => {
    try {
      await inventoryAPI.sellItem(itemId);
      await fetchInventory();
      return true;
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to sell item');
      console.error('Sell item error:', err);
      return false;
    }
  };

  return {
    inventory: data,
    loading,
    error,
    refresh: fetchInventory,
    sellItem,
  };
};
