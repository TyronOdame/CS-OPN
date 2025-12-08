'use client';

import { useState, useEffect } from 'react';
import { casesAPI } from '@/lib/api';
import { Case } from '@/lib/types';

// Custom hook to fetch and manage cases
export const userCases = () => {
  // state to hold the fetched cases
  const [cases, setCases] = useState<Case[]>([]);

  // state to track loading status
  const [loading, setLoading] = useState(true);

  // state to track any errors
  const [error, setError] = useState<string | null>(null);

  // userEffect to fetch cases on component mount
  useEffect(() => {
    const fetchCases = async () => {
      try {
        // 1. Set loading to true (shows spinner in UI)
        setLoading(true);

        // 2. Clear any previous errors
        setError(null);

        // 3. Call the API to get all cases
        const data = await casesAPI.getAllCases();

        // 4. Store the cases in state
        setCases(data);
      } catch (err) {
        // If API call fails, store the error message
        setError(err instanceof Error ? err.message : 'Failed to load cases');

        // Log error to console for debugging
        console.error('Cases fetch error:', err);
      } finally {
        // Always set loading to false when done (success or error)
        setLoading(false);
      }
    };

    // Execute the fetch function
    fetchCases();
  }, []);

  // Return the states so components can use them
  return { cases, loading, error };
};
