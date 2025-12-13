'use client';

import { useState, useEffect } from 'react';
import {
  Case,
  CaseOpenResult,
  RARITY_COLORS,
  RARITY_TEXT_COLORS,
} from '@/lib/types';
import { Dialog, DialogContent } from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Loader2, Sparkles, Package, DollarSign } from 'lucide-react';
import { caseAPI } from '@/lib/api';
import { on } from 'events';
import { set } from 'zod';

// this if the props for the caseOpeningModel

interface CaseOpeningModelProps {
  open: boolean;
  onClose: () => void;
  caseItem: Case;
  onSuccess: () => void;
}

/*
    caseOpeningModel component

    Handles the full case opening experience:
    1. Ready stage - Confirm opening
    2. Opening stage - Show animation
    3. Revealing stage - Animate skin reveal
    4. Revealed stage - Show final result
    
*/

export const CaseOpeningModal = ({
  open,
  onclose,
  caseItem,
  onSuccess,
}: CaseOpeningModelProps) => {
  // Track which stage of the opening we're in
  const [stage, setStage] = useState<
    'ready' | 'opening' | 'revealing' | 'revealed'
  >('ready');

  // Store the results from the API
  const [result, setResult] = useState<CaseOpenResult | null>(null);

  // Store any error messages
  const [error, setError] = useState<string | null>(null);

  // Reset state when modal opens/closes
  useEffect(() => {
    if (open) {
      setStage('ready');
      setResult(null);
      setError(null);
    }
  }, [open]);

  // Handle opening the case
  const handleOpenCase = async () => {
    setStage('opening');
    setError(null);

    try {
      //show spinning animation for 2 seconds
      await new Promise((resolve) => setTimeout(resolve, 2000));

      // call backend API to open the case
      const openResult = await caseAPI.openCase(caseItem.id);
      setResult(openResult);

      // show revealing stage
      setStage('revealing');

      // after 1 second, show revealed stage
      setTimeout(() => {
        setStage('revealed');
      }, 1000);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to open case');
      setStage('ready');
    }
  };
};
