/* eslint-disable @next/next/no-img-element */
'use client';

import { Case } from '@/lib/types';
import { Card, CardContent, CardFooter } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { DollarSign, Package } from 'lucide-react';

interface CaseCardProps {
  caseItem: Case;
  onBuy: (caseId: Case) => void;
  userBalance: number;
}

const FALLBACK_IMAGE_SRC = '/file.svg';

// Component to display individual case information
export const CaseCard = ({ caseItem, onBuy, userBalance }: CaseCardProps) => {
  // check if user has enough balance to open the case
  const canAfford = userBalance >= caseItem.price;

  return (
    <Card className="overflow-hidden transition-all hover:shadow-xl hover:scale-105 bg-gradient-to-br from-gray-900 to-gray-950 border-gray-800 group">
      {/* Case Image Section */}
      <div className="relative aspect-square bg-gradient-to-br from-gray-800 to-gray-900 overflow-hidden">
        <img
          src={caseItem.image_url}
          alt={caseItem.name}
          onError={(event) => {
            const img = event.currentTarget;
            if (img.dataset.fallbackApplied) return;
            img.dataset.fallbackApplied = 'true';
            img.src = FALLBACK_IMAGE_SRC;
          }}
          className="w-full h-full object-contain transition-transform group-hover:scale-110 p-8"
        />
        {/* Glow effect on hover */}
        <div className="absolute inset-0 bg-gradient-to-t from-blue-500/20 to-transparent opacity-0 group-hover:opacity-100 transition-opacity" />
      </div>

      {/* Case Details Section */}
      <CardContent className="p-6">
        {/* Case Name and Description */}
        <div className="flex items-start gap-2 mb-3">
          <Package className="w-5 h-5 text-blue-400 mt-1 flex-shrink-0" />
          <div>
            <h3 className="font-bold text-lg text-white mb-1">
              {caseItem.name}
            </h3>
            <p className="text-sm text-gray-400 line-clamp-2">
              {caseItem.description}
            </p>
          </div>
        </div>

        {/* Price Display */}
        <div className="flex items-center gap-2 text-green-400 font-bold text-xl mt-4">
          <DollarSign className="w-6 h-6" />
          <span>{caseItem.price.toFixed(2)} CB</span>
        </div>
      </CardContent>

      {/* Button Section */}
      <CardFooter className="p-6 pt-0">
        <Button
          onClick={() => onBuy(caseItem)}
          disabled={!canAfford}
          className="w-full bg-blue-600 hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {canAfford ? 'Buy Case' : 'Insufficient Funds'}
        </Button>
      </CardFooter>
    </Card>
  );
};
