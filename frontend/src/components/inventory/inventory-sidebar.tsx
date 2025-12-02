'use client';

import { Card } from '@/components/ui/card';
import { ChevronDown } from 'lucide-react';
import { useState } from 'react';

interface InventorySidebarProps {
  stats?: Record<string, number>;
  sortBy: 'recent' | 'value' | 'rarity';
  setSortBy: (sort: 'recent' | 'value' | 'rarity') => void;
  filterRarity: string | null;
  setFilterRarity: (rarity: string | null) => void;
}

const RARITY_ORDER = [
  'Covert',
  'Classified',
  'Restricted',
  'Mil-Spec',
  'Industrial Grade',
  'Consumer Grade',
];

const RARITY_COLORS: Record<string, string> = {
  Covert: 'bg-red-600/20 text-red-400 border-red-600/50',
  Classified: 'bg-purple-600/20 text-purple-400 border-purple-600/50',
  Restricted: 'bg-blue-600/20 text-blue-400 border-blue-600/50',
  'Mil-Spec': 'bg-cyan-600/20 text-cyan-400 border-cyan-600/50',
  'Industrial Grade': 'bg-amber-600/20 text-amber-400 border-amber-600/50',
  'Consumer Grade': 'bg-gray-600/20 text-gray-400 border-gray-600/50',
};

export default function InventorySidebar({
  stats,
  sortBy,
  setSortBy,
  filterRarity,
  setFilterRarity,
}: InventorySidebarProps) {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div className="w-full lg:w-64 space-y-4 md:w-auto">
      {/* Sort Section */}
      <div className="md:flex md:flex-row md:gap-3 md:items-center">
        {/* Sort dropdown */}
        <div className="relative mb-3 md:mb-0">
          <button
            onClick={() => setIsOpen(!isOpen)}
            className="flex items-center gap-2 px-3 py-2 bg-card hover:bg-card/80 border border-border rounded text-sm text-foreground transition-all w-full md:w-auto"
          >
            <span>
              {sortBy === 'recent' && 'Recently Acquired'}
              {sortBy === 'value' && 'Highest Value'}
              {sortBy === 'rarity' && 'Rarity'}
            </span>
            <ChevronDown className="w-4 h-4" />
          </button>

          {isOpen && (
            <div className="absolute top-full left-0 mt-1 bg-card border border-border rounded shadow-lg z-10 w-full md:w-48">
              {(['recent', 'value', 'rarity'] as const).map((sort) => (
                <button
                  key={sort}
                  onClick={() => {
                    setSortBy(sort);
                    setIsOpen(false);
                  }}
                  className={`w-full text-left px-3 py-2 text-sm hover:bg-primary/20 transition-all ${
                    sortBy === sort
                      ? 'bg-primary/30 text-primary'
                      : 'text-foreground'
                  }`}
                >
                  {sort === 'recent' && 'Recently Acquired'}
                  {sort === 'value' && 'Highest Value'}
                  {sort === 'rarity' && 'Rarity'}
                </button>
              ))}
            </div>
          )}
        </div>

        {/* Filter Section */}
        <div className="flex flex-wrap gap-1">
          {RARITY_ORDER.map((rarity) => {
            const count = stats?.[rarity] || 0;
            const isSelected = filterRarity === rarity;

            return (
              <button
                key={rarity}
                onClick={() => setFilterRarity(isSelected ? null : rarity)}
                className={`px-2 py-1 text-xs font-medium rounded transition-all ${
                  isSelected
                    ? `${RARITY_COLORS[rarity]} bg-black/40 border border-current`
                    : 'text-muted-foreground bg-card/50 border border-transparent hover:bg-card'
                }`}
              >
                {rarity.split(' ')[0]}
                {count > 0 && <span className="ml-1">({count})</span>}
              </button>
            );
          })}
        </div>
      </div>

      {/* Stats Section */}
    </div>
  );
}
