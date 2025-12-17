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
import { casesAPI } from '@/lib/api';

/**
 * Props for CaseOpeningModal
 */
interface CaseOpeningModalProps {
  open: boolean; // Is modal visible?
  onClose: () => void; // Function to close modal
  caseItem: Case; // The case being opened
  onSuccess: () => void; // Function to call after successful opening (refresh inventory)
}

/**
 * CaseOpeningModal Component
 *
 * Handles the full case opening experience:
 * 1. Ready stage - Confirm opening
 * 2. Opening stage - Show animation
 * 3. Revealing stage - Animate skin reveal
 * 4. Revealed stage - Show final result
 */
export const CaseOpeningModal = ({
  open,
  onClose,
  caseItem,
  onSuccess,
}: CaseOpeningModalProps) => {
  // Track which stage of the opening we're in
  const [stage, setStage] = useState<
    'ready' | 'opening' | 'revealing' | 'revealed'
  >('ready');

  // Store the result from the API
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

  /**
   * Handle opening the case
   * This function:
   * 1. Shows opening animation
   * 2. Calls backend API
   * 3. Shows revealing animation
   * 4. Shows final result
   */
  const handleOpenCase = async () => {
    setStage('opening');
    setError(null);

    try {
      // Show spinning animation for 2 seconds
      await new Promise((resolve) => setTimeout(resolve, 2000));

      // Call backend API to open the case
      const openResult = await casesAPI.openCase(caseItem.id);
      setResult(openResult);

      // Show revealing stage (skin bounces in)
      setStage('revealing');

      // After 1 second, show final revealed state
      setTimeout(() => {
        setStage('revealed');
      }, 1000);
    } catch (err) {
      // If something goes wrong, show error and go back to ready
      setError(err instanceof Error ? err.message : 'Failed to open case');
      setStage('ready');
    }
  };

  /**
   * Handle closing the modal
   * If we successfully opened a case, refresh the inventory
   */
  const handleClose = () => {
    if (stage === 'revealed') {
      onSuccess(); // Refresh inventory/balance
    }
    onClose(); // Close modal
  };

  // Get the color gradient for the skin rarity
  const rarityGradient = result
    ? RARITY_COLORS[result.skin.rarity] || 'from-gray-400 to-gray-500'
    : 'from-blue-400 to-blue-500';

  // Get the text color for the skin rarity
  const rarityTextColor = result
    ? RARITY_TEXT_COLORS[result.skin.rarity] || 'text-gray-400'
    : 'text-blue-400';

  return (
    <Dialog open={open} onOpenChange={handleClose}>
      <DialogContent className="sm:max-w-2xl bg-gray-950 border-gray-800 text-white">
        {/* ========================================
            STAGE 1: READY - Confirm Opening
            ======================================== */}
        {stage === 'ready' && (
          <div className="text-center py-8">
            {/* Case Image */}
            <div className="mb-6">
              <img
                src={caseItem.image_url}
                alt={caseItem.name}
                className="w-48 h-48 mx-auto object-contain"
              />
            </div>

            {/* Case Name */}
            <h2 className="text-3xl font-bold mb-2">{caseItem.name}</h2>

            {/* Case Description */}
            <p className="text-gray-400 mb-6">{caseItem.description}</p>

            {/* Price */}
            <div className="flex items-center justify-center gap-2 text-green-400 text-2xl font-bold mb-8">
              <DollarSign className="w-6 h-6" />
              <span>{caseItem.price.toFixed(2)} CB</span>
            </div>

            {/* Error Message */}
            {error && (
              <div className="bg-red-900/20 border border-red-500/50 rounded-lg p-4 text-red-400 text-sm mb-4">
                {error}
              </div>
            )}

            {/* Action Buttons */}
            <div className="flex gap-4">
              <Button
                variant="outline"
                onClick={handleClose}
                className="flex-1"
              >
                Cancel
              </Button>
              <Button
                onClick={handleOpenCase}
                className="flex-1 bg-blue-600 hover:bg-blue-700"
              >
                <Package className="w-4 h-4 mr-2" />
                Open Case
              </Button>
            </div>
          </div>
        )}

        {/* ========================================
            STAGE 2: OPENING - Spinning Animation
            ======================================== */}
        {stage === 'opening' && (
          <div className="text-center py-16">
            {/* Spinning Loader */}
            <div className="relative mb-8">
              <Loader2 className="w-24 h-24 mx-auto animate-spin text-blue-500" />
              <Sparkles className="w-12 h-12 absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 text-yellow-400 animate-pulse" />
            </div>

            {/* Loading Text */}
            <h2 className="text-2xl font-bold mb-2">Opening Case...</h2>
            <p className="text-gray-400">Preparing your reward</p>
          </div>
        )}

        {/* ========================================
            STAGE 3 & 4: REVEALING / REVEALED
            ======================================== */}
        {(stage === 'revealing' || stage === 'revealed') && result && (
          <div className="text-center py-8">
            {/* Rarity Color Stripe */}
            <div
              className={`h-2 bg-gradient-to-r ${rarityGradient} mb-6 rounded-full animate-pulse`}
            />

            {/* Skin Image */}
            <div
              className={`relative mb-6 ${
                stage === 'revealing' ? 'animate-bounce' : ''
              }`}
            >
              <div className="relative inline-block">
                <img
                  src={result.skin.image_url}
                  alt={result.skin.name}
                  className="w-64 h-64 object-contain mx-auto"
                />
                {/* Glow Effect (only when fully revealed) */}
                {stage === 'revealed' && (
                  <div className="absolute inset-0 bg-gradient-to-t from-transparent via-transparent to-yellow-400/30 animate-pulse" />
                )}
              </div>
            </div>

            {/* Skin Details */}
            <div className="space-y-3">
              {/* Skin Name */}
              <h2 className="text-3xl font-bold">{result.skin.name}</h2>

              {/* Rarity */}
              <div className={`text-xl font-semibold ${rarityTextColor}`}>
                {result.skin.rarity}
              </div>

              {/* Condition */}
              <div className="text-lg text-gray-300">
                {result.inventory_item.condition}
              </div>

              {/* Float Value */}
              <div className="flex items-center justify-center gap-2 text-gray-400 text-sm">
                <span>Float: {result.inventory_item.float.toFixed(4)}</span>
              </div>

              {/* Item Value */}
              <div className="flex items-center justify-center gap-2 text-green-400 text-2xl font-bold mt-4">
                <DollarSign className="w-6 h-6" />
                <span>{result.inventory_item.value.toFixed(2)} CB</span>
              </div>

              {/* New Balance */}
              <div className="text-sm text-gray-500 mt-2">
                New Balance: ${result.new_balance.toFixed(2)} CB
              </div>
            </div>

            {/* Action Buttons (only when fully revealed) */}
            {stage === 'revealed' && (
              <div className="mt-8 space-y-3">
                <Button
                  onClick={handleClose}
                  className="w-full bg-green-600 hover:bg-green-700"
                >
                  <Sparkles className="w-4 h-4 mr-2" />
                  Awesome! Add to Inventory
                </Button>
                <Button
                  variant="outline"
                  onClick={() => {
                    setStage('ready');
                    setResult(null);
                  }}
                  className="w-full"
                >
                  Open Another Case
                </Button>
              </div>
            )}
          </div>
        )}
      </DialogContent>
    </Dialog>
  );
};
