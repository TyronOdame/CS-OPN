'use client';

import { useState, useEffect } from 'react';
import Link from 'next/link';
import { Card } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import SkinCard from '@/components/inventory/skin-card';
import InventoryHeader from '@/components/inventory/inventory-header';
import InventorySidebar from '@/components/inventory/inventory-sidebar';

interface Skin {
  id: string;
  name: string;
  rarity: string;
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

interface InventoryData {
  items: InventoryItem[];
  total_value: number;
  item_count: number;
  stats: Record<string, number>;
}

const DEMO_DATA: InventoryData = {
  items: [
    {
      id: '1',
      skin: {
        id: 's1',
        name: 'Dragon Lore',
        rarity: 'Covert',
        image_url: '/cs2-dragon-lore-skin.jpg',
      },
      float: 0.0234,
      condition: 'Factory New',
      value: 450.0,
      acquired_from: 'Souvenir',
      acquired_at: new Date(Date.now() - 2 * 24 * 60 * 60 * 1000).toISOString(),
    },
    {
      id: '2',
      skin: {
        id: 's2',
        name: 'Fade',
        rarity: 'Covert',
        image_url: '/cs2-fade-skin.jpg',
      },
      float: 0.0567,
      condition: 'Minimal Wear',
      value: 320.0,
      acquired_from: 'Case',
      acquired_at: new Date(Date.now() - 5 * 24 * 60 * 60 * 1000).toISOString(),
    },
    {
      id: '3',
      skin: {
        id: 's3',
        name: 'Crimson Web',
        rarity: 'Classified',
        image_url: '/cs2-crimson-web-skin.jpg',
      },
      float: 0.1234,
      condition: 'Field-Tested',
      value: 210.0,
      acquired_from: 'Case',
      acquired_at: new Date(Date.now() - 7 * 24 * 60 * 60 * 1000).toISOString(),
    },
    {
      id: '4',
      skin: {
        id: 's4',
        name: 'Phantom Disruptor',
        rarity: 'Classified',
        image_url: '/cs2-phantom-disruptor-skin.jpg',
      },
      float: 0.2345,
      condition: 'Well-Worn',
      value: 95.0,
      acquired_from: 'Case',
      acquired_at: new Date(
        Date.now() - 10 * 24 * 60 * 60 * 1000
      ).toISOString(),
    },
    {
      id: '5',
      skin: {
        id: 's5',
        name: 'Point Disarray',
        rarity: 'Restricted',
        image_url: '/cs2-point-disarray-skin.jpg',
      },
      float: 0.3456,
      condition: 'Minimal Wear',
      value: 78.0,
      acquired_from: 'Case',
      acquired_at: new Date(
        Date.now() - 12 * 24 * 60 * 60 * 1000
      ).toISOString(),
    },
    {
      id: '6',
      skin: {
        id: 's6',
        name: 'Hyper Beast',
        rarity: 'Restricted',
        image_url: '/cs2-hyper-beast-skin.jpg',
      },
      float: 0.1567,
      condition: 'Factory New',
      value: 125.0,
      acquired_from: 'Case',
      acquired_at: new Date(Date.now() - 3 * 24 * 60 * 60 * 1000).toISOString(),
    },
    {
      id: '7',
      skin: {
        id: 's7',
        name: 'Rat Rod',
        rarity: 'Mil-Spec',
        image_url: '/cs2-rat-rod-skin.jpg',
      },
      float: 0.0891,
      condition: 'Minimal Wear',
      value: 12.5,
      acquired_from: 'Case',
      acquired_at: new Date(
        Date.now() - 15 * 24 * 60 * 60 * 1000
      ).toISOString(),
    },
    {
      id: '8',
      skin: {
        id: 's8',
        name: 'Anodized Navy',
        rarity: 'Industrial Grade',
        image_url: '/cs2-anodized-navy-skin.jpg',
      },
      float: 0.2134,
      condition: 'Field-Tested',
      value: 2.5,
      acquired_from: 'Case',
      acquired_at: new Date(
        Date.now() - 20 * 24 * 60 * 60 * 1000
      ).toISOString(),
    },
    {
      id: '9',
      skin: {
        id: 's9',
        name: 'Predator',
        rarity: 'Consumer Grade',
        image_url: '/cs2-predator-skin.jpg',
      },
      float: 0.4567,
      condition: 'Battle-Scarred',
      value: 0.5,
      acquired_from: 'Case',
      acquired_at: new Date(
        Date.now() - 25 * 24 * 60 * 60 * 1000
      ).toISOString(),
    },
  ],
  total_value: 1293.5,
  item_count: 9,
  stats: {
    Covert: 2,
    Classified: 2,
    Restricted: 2,
    'Mil-Spec': 1,
    'Industrial Grade': 1,
    'Consumer Grade': 1,
  },
};

export default function InventoryPage() {
  const [inventory, setInventory] = useState<InventoryData | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [sortBy, setSortBy] = useState<'recent' | 'value' | 'rarity'>('recent');
  const [filterRarity, setFilterRarity] = useState<string | null>(null);
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    const fetchInventory = async () => {
      try {
        setLoading(true);
        const token = localStorage.getItem('token');
        setIsLoggedIn(!!token);

        if (!token) {
          setInventory(DEMO_DATA);
          setError(null);
          setLoading(false);
          return;
        }

        const response = await fetch('/api/inventory', {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        if (!response.ok) {
          throw new Error('Failed to fetch inventory');
        }

        const data = await response.json();
        setInventory(data);
        setError(null);
      } catch (err) {
        setError(
          err instanceof Error ? err.message : 'Failed to load inventory'
        );
        setInventory(null);
      } finally {
        setLoading(false);
      }
    };

    fetchInventory();
  }, []);

  const sortedItems = inventory?.items
    ? [...inventory.items].sort((a, b) => {
        switch (sortBy) {
          case 'value':
            return b.value - a.value;
          case 'rarity':
            const rarityOrder = [
              'Covert',
              'Classified',
              'Restricted',
              'Mil-Spec',
              'Industrial Grade',
              'Consumer Grade',
            ];
            return (
              rarityOrder.indexOf(b.skin.rarity) -
              rarityOrder.indexOf(a.skin.rarity)
            );
          case 'recent':
          default:
            return (
              new Date(b.acquired_at).getTime() -
              new Date(a.acquired_at).getTime()
            );
        }
      })
    : [];

  const filteredItems = filterRarity
    ? sortedItems.filter((item) => item.skin.rarity === filterRarity)
    : sortedItems;

  if (loading) {
    return (
      <div className="min-h-screen bg-background">
        <InventoryHeader />
        <div className="flex items-center justify-center h-96">
          <p className="text-muted-foreground text-lg">
            Loading your inventory...
          </p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="min-h-screen bg-background">
        <InventoryHeader />
        <div className="flex items-center justify-center h-96">
          <Card className="p-8 text-center max-w-md">
            <p className="text-destructive mb-4">{error}</p>
            <Link href="/login">
              <Button variant="default">Go to Login</Button>
            </Link>
          </Card>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-background">
      <InventoryHeader />

      <div className="container mx-auto px-4 py-8 text-muted-foreground bg-secondary my-10">
        {!isLoggedIn && (
          <div className="mb-6 p-3 bg-blue-600/20 border border-blue-600/50 rounded-md text-blue-400 text-sm">
            Viewing demo inventory. Log in to see your real inventory!
          </div>
        )}

        <div className="flex gap-6">
          <div className="w-48 flex-shrink-0">
            <Card className="p-4 bg-secondary/50 border-secondary">
              <div className="mb-6">
                <p className="text-xs text-muted-foreground uppercase tracking-wider mb-3 font-semibold">
                  Total Value
                </p>
                <p className="text-2xl font-bold text-foreground">
                  ${(inventory?.total_value || 0).toFixed(2)}
                </p>
                <p className="text-xs text-muted-foreground mt-1">
                  {filteredItems.length} items
                </p>
              </div>

              <div className="border-t border-secondary-foreground/20 pt-4">
                <p className="text-xs text-muted-foreground uppercase tracking-wider mb-3 font-semibold">
                  Rarity Breakdown
                </p>
                <div className="space-y-2">
                  {inventory?.stats &&
                    Object.entries(inventory.stats).map(([rarity, count]) => (
                      <div
                        key={rarity}
                        className="flex items-center justify-between text-xs"
                      >
                        <span className="text-muted-foreground">{rarity}</span>
                        <span className="font-semibold text-foreground">
                          {count}
                        </span>
                      </div>
                    ))}
                </div>
              </div>
            </Card>
          </div>

          <div className="flex-1">
            {filteredItems.length === 0 ? (
              <Card className="p-12 text-center">
                <p className="text-muted-foreground mb-4">
                  No skins in your inventory yet
                </p>
                <Link href="/">
                  <Button>Open Cases</Button>
                </Link>
              </Card>
            ) : (
              <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-3">
                {filteredItems.map((item) => (
                  <Link key={item.id} href={`/inventory/${item.id}`}>
                    <SkinCard item={item} />
                  </Link>
                ))}
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
