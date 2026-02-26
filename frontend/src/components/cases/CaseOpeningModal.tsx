'use client';

import { useEffect, useRef, useState } from 'react';
import {
  Case,
  CaseDetails,
  CaseOpenResult,
  RARITY_COLORS,
  RARITY_TEXT_COLORS,
  Skin,
} from '@/lib/types';
import { Dialog, DialogContent } from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Sparkles, Package, DollarSign } from 'lucide-react';
import { casesAPI, inventoryAPI } from '@/lib/api';

const FALLBACK_IMAGE_SRC = '/file.svg';

/**
 * Props for CaseOpeningModal
 */
interface CaseOpeningModalProps {
  open: boolean; // Is modal visible?
  onClose: () => void; // Function to close modal
  caseItem: Case; // The case being opened
  onSuccess: () => void; // Function to call after successful opening (refresh inventory)
  purchasedCaseId: string;
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
  purchasedCaseId,
}: CaseOpeningModalProps) => {
  const [stage, setStage] = useState<'ready' | 'opening' | 'revealed'>('ready');
  const [result, setResult] = useState<CaseOpenResult | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [caseDetails, setCaseDetails] = useState<CaseDetails | null>(null);
  const [loadingCaseDetails, setLoadingCaseDetails] = useState(false);
  const [reelSkins, setReelSkins] = useState<Skin[]>([]);
  const [reelOffset, setReelOffset] = useState(0);
  const [isRolling, setIsRolling] = useState(false);
  const reelViewportRef = useRef<HTMLDivElement>(null);

  const ROLL_ITEM_WIDTH = 140;
  const ROLL_GAP = 12;
  const MAX_VISIBLE_WIDTH = 720;
  const ROLL_DURATION_MS = 4500;

  useEffect(() => {
    if (open) {
      setStage('ready');
      setResult(null);
      setError(null);
      setReelSkins([]);
      setReelOffset(0);
      setIsRolling(false);

      const loadCaseDetails = async () => {
        try {
          setLoadingCaseDetails(true);
          const details = await casesAPI.getCaseById(caseItem.id);
          setCaseDetails(details);
        } catch (err) {
          setError(
            err instanceof Error ? err.message : 'Failed to load case contents'
          );
        } finally {
          setLoadingCaseDetails(false);
        }
      };

      loadCaseDetails();
    }
  }, [open, caseItem.id]);

  const pickWeightedSkin = (pool: CaseDetails['skins']): Skin => {
    if (!pool.length) {
      throw new Error('No skins available in this case');
    }
    const totalWeight = pool.reduce(
      (sum, skin) => sum + (skin.drop_chance || 0),
      0
    );
    if (totalWeight <= 0) {
      return pool[Math.floor(Math.random() * pool.length)];
    }
    const roll = Math.random() * totalWeight;
    let cumulative = 0;
    for (const skin of pool) {
      cumulative += skin.drop_chance || 0;
      if (roll <= cumulative) return skin;
    }
    return pool[pool.length - 1];
  };

  const handleOpenCase = async () => {
    if (loadingCaseDetails) return;

    setStage('opening');
    setError(null);

    try {
      const openResult = await inventoryAPI.openPurchasedCase(purchasedCaseId);
      const pool = caseDetails?.skins || [];
      const winnerIndex = 42;
      const totalItems = 56;
      const generated: Skin[] = [];

      for (let i = 0; i < totalItems; i++) {
        generated.push(
          pool.length ? pickWeightedSkin(pool) : { ...openResult.skin }
        );
      }
      generated[winnerIndex] = openResult.skin;

      setIsRolling(false);
      setReelSkins(generated);
      setReelOffset(0);
      setResult(openResult);

      requestAnimationFrame(() => {
        const currentViewportWidth =
          reelViewportRef.current?.clientWidth || MAX_VISIBLE_WIDTH;
        const centerOffset = currentViewportWidth / 2 - ROLL_ITEM_WIDTH / 2;
        const targetOffset =
          -(winnerIndex * (ROLL_ITEM_WIDTH + ROLL_GAP) - centerOffset);
        setIsRolling(true);
        setReelOffset(targetOffset);
      });

      setTimeout(() => {
        setStage('revealed');
      }, ROLL_DURATION_MS + 150);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to open case');
      setStage('ready');
    }
  };

  const handleClose = () => {
    if (stage === 'revealed') {
      onSuccess();
    }
    onClose();
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
                onError={(event) => {
                  const img = event.currentTarget;
                  if (img.dataset.fallbackApplied) return;
                  img.dataset.fallbackApplied = 'true';
                  img.src = FALLBACK_IMAGE_SRC;
                }}
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
                disabled={loadingCaseDetails}
                className="flex-1 bg-blue-600 hover:bg-blue-700"
              >
                <Package className="w-4 h-4 mr-2" />
                {loadingCaseDetails ? 'Loading...' : 'Open Case'}
              </Button>
            </div>
          </div>
        )}

        {/* Opening animation */}
        {stage === 'opening' && (
          <div className="py-8">
            <h2 className="mb-2 text-center text-2xl font-bold">Opening Case...</h2>
            <p className="mb-6 text-center text-gray-400">
              Watch the reel spin and stop on your item
            </p>

            <div className="px-4 sm:px-6">
              <div ref={reelViewportRef} className="relative mx-auto w-full max-w-[720px]">
                <div className="relative overflow-hidden rounded-lg border border-gray-800 bg-gray-900/80 py-4">
                  <div
                    className="flex"
                    style={{
                      gap: `${ROLL_GAP}px`,
                      transform: `translateX(${reelOffset}px)`,
                      transition: isRolling
                        ? `transform ${ROLL_DURATION_MS}ms cubic-bezier(0.12, 0.75, 0.2, 1)`
                        : 'none',
                    }}
                  >
                    {reelSkins.map((skin, index) => {
                      const rarityGradient =
                        RARITY_COLORS[skin.rarity] || 'from-gray-500 to-gray-600';
                      return (
                        <div
                          key={`${skin.id}-${index}`}
                          className="shrink-0 rounded-lg border border-gray-700 bg-gray-950 p-2"
                          style={{ width: `${ROLL_ITEM_WIDTH}px` }}
                        >
                          <div
                            className={`mb-2 h-1 rounded bg-gradient-to-r ${rarityGradient}`}
                          />
                          <img
                            src={skin.image_url}
                            alt={skin.name}
                            onError={(event) => {
                              const img = event.currentTarget;
                              if (img.dataset.fallbackApplied) return;
                              img.dataset.fallbackApplied = 'true';
                              img.src = FALLBACK_IMAGE_SRC;
                            }}
                            className="mx-auto h-20 w-20 object-contain"
                          />
                          <p className="line-clamp-2 text-center text-xs text-gray-200">
                            {skin.name}
                          </p>
                        </div>
                      );
                    })}
                  </div>
                  <div className="pointer-events-none absolute inset-y-0 left-1/2 z-30 w-[2px] -translate-x-1/2 bg-yellow-400 shadow-[0_0_12px_rgba(250,204,21,0.95)]" />
                </div>
              </div>
            </div>
          </div>
        )}

        {/* Final reveal */}
        {stage === 'revealed' && result && (
          <div className="text-center py-8">
            <div
              className={`h-2 bg-gradient-to-r ${rarityGradient} mb-6 rounded-full animate-pulse`}
            />

            <div className="relative mb-6">
              <div className="relative inline-block">
                <img
                  src={result.skin.image_url}
                  alt={result.skin.name}
                  onError={(event) => {
                    const img = event.currentTarget;
                    if (img.dataset.fallbackApplied) return;
                    img.dataset.fallbackApplied = 'true';
                    img.src = FALLBACK_IMAGE_SRC;
                  }}
                  className="w-64 h-64 object-contain mx-auto"
                />
                <div className="absolute inset-0 bg-gradient-to-t from-transparent via-transparent to-yellow-400/30 animate-pulse" />
              </div>
            </div>

            <div className="space-y-3">
              <h2 className="text-3xl font-bold">{result.skin.name}</h2>
              <div className={`text-xl font-semibold ${rarityTextColor}`}>
                {result.skin.rarity}
              </div>
              <div className="text-lg text-gray-300">{result.condition}</div>
              <div className="flex items-center justify-center gap-2 text-gray-400 text-sm">
                <span>Float: {result.float.toFixed(4)}</span>
              </div>
              <div className="flex items-center justify-center gap-2 text-green-400 text-2xl font-bold mt-4">
                <DollarSign className="w-6 h-6" />
                <span>{result.value.toFixed(2)} CB</span>
              </div>
              <div className="text-sm text-gray-500 mt-2">
                New Balance: ${result.new_balance.toFixed(2)} CB
              </div>
            </div>

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
                  setError(null);
                }}
                className="w-full"
              >
                Open Another Case
              </Button>
            </div>
          </div>
        )}
      </DialogContent>
    </Dialog>
  );
};
