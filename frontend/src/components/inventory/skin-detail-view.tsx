'use client';

import { Card } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import Image from 'next/image';
import { useState } from 'react';
import { inventoryAPI } from '@/lib/api';

interface Skin {
  id: string;
  name: string;
  rarity: string;
  min_value: number;
  max_value: number;
  image_url?: string;
}

interface InventoryItem {
  id: string;
  skin: Skin;
  float: number;
  condition: string;
  value: number;
  acquired_from: string;
  acquired_at: string;
}

interface SkinDetailViewProps {
  item: InventoryItem;
}

const FALLBACK_IMAGE_SRC = '/file.svg';

const RARITY_INFO: Record<string, { color: string; description: string }> = {
  Covert: { color: 'text-red-400', description: 'Extremely Rare' },
  Classified: { color: 'text-purple-400', description: 'Very Rare' },
  Restricted: { color: 'text-blue-400', description: 'Rare' },
  'Mil-Spec': { color: 'text-cyan-400', description: 'Uncommon' },
  'Industrial Grade': { color: 'text-amber-400', description: 'Common' },
  'Consumer Grade': { color: 'text-gray-400', description: 'Very Common' },
};

const CONDITION_INFO: Record<string, { range: string; description: string }> = {
  'Factory New': { range: '0.00 - 0.07', description: 'Looks freshly unboxed' },
  'Minimal Wear': { range: '0.07 - 0.14', description: 'Looks barely used' },
  'Field-Tested': { range: '0.14 - 0.38', description: 'Shows signs of use' },
  'Well-Worn': {
    range: '0.38 - 0.45',
    description: 'Clearly used in the field',
  },
  'Battle-Scarred': {
    range: '0.45 - 1.00',
    description: 'Heavily used in combat',
  },
};

export default function SkinDetailView({ item }: SkinDetailViewProps) {
  const [showSellModal, setShowSellModal] = useState(false);
  const [selling, setSelling] = useState(false);

  const rarityInfo =
    RARITY_INFO[item.skin.rarity] || RARITY_INFO['Consumer Grade'];
  const conditionInfo = CONDITION_INFO[item.condition];

  const handleSell = async () => {
    try {
      setSelling(true);
      await inventoryAPI.sellItem(item.id);

      // Redirect to inventory after successful sale
      window.location.href = '/inventory';
    } catch {
      alert('Failed to sell skin');
      setSelling(false);
    }
  };

  return (
    <main className="container mx-auto px-4 py-8">
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        {/* Main Image */}
        <div className="lg:col-span-2">
          <Card className="overflow-hidden border-2 border-primary/20 bg-card/50">
            <div className="relative w-full bg-gradient-to-br from-card to-background p-8 aspect-square flex items-center justify-center">
              {item.skin.image_url ? (
                <Image
                  src={item.skin.image_url || FALLBACK_IMAGE_SRC}
                  alt={item.skin.name}
                  width={500}
                  height={500}
                  unoptimized
                  onError={(event) => {
                    const img = event.currentTarget as HTMLImageElement;
                    if (img.dataset.fallbackApplied) return;
                    img.dataset.fallbackApplied = 'true';
                    img.src = FALLBACK_IMAGE_SRC;
                  }}
                  className="max-w-full max-h-full object-contain"
                  priority
                />
              ) : (
                <div className="text-center">
                  <p className="text-muted-foreground text-lg">Skin Preview</p>
                  <p className="text-muted-foreground text-sm mt-2">
                    {item.skin.name}
                  </p>
                </div>
              )}
            </div>
          </Card>

          {/* Details Grid */}
          <div className="grid grid-cols-2 gap-4 mt-8">
            <Card className="p-4">
              <p className="text-xs text-muted-foreground mb-2">Condition</p>
              <p className="text-2xl font-bold text-foreground">
                {item.condition}
              </p>
              {conditionInfo && (
                <p className="text-xs text-muted-foreground mt-2">
                  {conditionInfo.description}
                </p>
              )}
            </Card>
            <Card className="p-4">
              <p className="text-xs text-muted-foreground mb-2">Float Value</p>
              <p className="text-2xl font-mono font-bold text-primary">
                {item.float.toFixed(4)}
              </p>
              {conditionInfo && (
                <p className="text-xs text-muted-foreground mt-2">
                  Range: {conditionInfo.range}
                </p>
              )}
            </Card>
          </div>
        </div>

        {/* Sidebar */}
        <div className="space-y-4">
          {/* Title and Rarity */}
          <Card className="p-6 bg-gradient-to-br from-card to-background">
            <h1 className="text-2xl md:text-3xl font-bold text-foreground mb-2">
              {item.skin.name}
            </h1>
            <Badge
              variant="outline"
              className={`${rarityInfo.color} text-lg px-4 py-2`}
            >
              {item.skin.rarity} - {rarityInfo.description}
            </Badge>
          </Card>

          {/* Value Information */}
          <Card className="p-6 border-accent/20">
            <p className="text-sm text-muted-foreground mb-2">Current Value</p>
            <p className="text-4xl font-bold text-accent mb-2">
              ${item.value.toFixed(2)}
            </p>
            <div className="text-xs text-muted-foreground space-y-1">
              <p>Min: ${item.skin.min_value.toFixed(2)}</p>
              <p>Max: ${item.skin.max_value.toFixed(2)}</p>
            </div>
          </Card>

          {/* Acquisition Info */}
          <Card className="p-6">
            <div className="space-y-3">
              <div>
                <p className="text-xs text-muted-foreground mb-1">
                  Obtained From
                </p>
                <p className="text-foreground font-semibold">
                  {item.acquired_from}
                </p>
              </div>
              <div>
                <p className="text-xs text-muted-foreground mb-1">
                  Date Acquired
                </p>
                <p className="text-foreground font-semibold">
                  {new Date(item.acquired_at).toLocaleDateString()}
                </p>
              </div>
            </div>
          </Card>

          {/* Actions */}
          <div className="space-y-2 pt-4">
            <Button
              className="w-full cs2-glow"
              size="lg"
              onClick={() => setShowSellModal(true)}
            >
              Sell Skin
            </Button>
            <p className="text-xs text-muted-foreground text-center">
              Sell this skin to receive ${item.value.toFixed(2)} in Case Bucks
            </p>
          </div>

          {/* Sell Confirmation Modal */}
          {showSellModal && (
            <Card className="p-6 border-destructive/50 bg-destructive/10">
              <p className="text-sm text-foreground mb-4">
                Are you sure you want to sell this skin for{' '}
                <span className="font-bold text-primary">
                  ${item.value.toFixed(2)}
                </span>
                ?
              </p>
              <div className="flex gap-2">
                <Button
                  variant="destructive"
                  className="flex-1"
                  disabled={selling}
                  onClick={handleSell}
                >
                  {selling ? 'Selling...' : 'Confirm Sell'}
                </Button>
                <Button
                  variant="outline"
                  className="flex-1 bg-transparent"
                  disabled={selling}
                  onClick={() => setShowSellModal(false)}
                >
                  Cancel
                </Button>
              </div>
            </Card>
          )}
        </div>
      </div>
    </main>
  );
}
