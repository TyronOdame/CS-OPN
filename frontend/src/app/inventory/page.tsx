'use client';

import { useEffect, useMemo, useState } from 'react';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { Card } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import SkinCard from '@/components/inventory/skin-card';
import InventoryHeader from '@/components/inventory/inventory-header';
import { useInventory } from '@/hooks/useInventory';
import { authAPI, inventoryAPI, aiAPI } from '@/lib/api';
import { Loader2, Package, Sparkles, ArrowLeft } from 'lucide-react';
import { CaseOpeningModal } from '@/components/cases/CaseOpeningModal';
import { PurchasedCase } from '@/lib/types';

export default function InventoryPage() {
  const router = useRouter();
  const { inventory, loading, error, refresh } = useInventory(false);
  const [sortBy, setSortBy] = useState<'recent' | 'value' | 'rarity'>('recent');
  const [filterRarity, setFilterRarity] = useState<string | null>(null);
  const [mounted, setMounted] = useState(false);
  const [purchasedCases, setPurchasedCases] = useState<PurchasedCase[]>([]);
  const [loadingPurchasedCases, setLoadingPurchasedCases] = useState(true);
  const [selectedPurchasedCase, setSelectedPurchasedCase] =
    useState<PurchasedCase | null>(null);
  const [priceInput, setPriceInput] = useState('');
  const [priceResult, setPriceResult] = useState<string | null>(null);
  const [priceLoading, setPriceLoading] = useState(false);

  useEffect(() => {
    setMounted(true);
    if (!authAPI.isAuthenticated()) {
      router.push('/login');
      return;
    }
    const fetchPurchasedCases = async () => {
      try {
        setLoadingPurchasedCases(true);
        const data = await inventoryAPI.getPurchasedCases();
        setPurchasedCases(data.cases || []);
      } catch (err) {
        console.error('Failed to fetch purchased cases:', err);
      } finally {
        setLoadingPurchasedCases(false);
      }
    };
    fetchPurchasedCases();
  }, [router]);

  const refreshAll = async () => {
    await refresh();
    const data = await inventoryAPI.getPurchasedCases();
    setPurchasedCases(data.cases || []);
  };

  const sortedItems = useMemo(() => {
    const items = inventory?.items ? [...inventory.items] : [];
    return items.sort((a, b) => {
      switch (sortBy) {
        case 'value':
          return b.value - a.value;
        case 'rarity': {
          const rarityOrder = [
            'Rare Special',
            'Covert',
            'Classified',
            'Restricted',
            'Mil-Spec',
            'Industrial Grade',
            'Consumer Grade',
          ];
          return (
            rarityOrder.indexOf(b.skin.rarity) - rarityOrder.indexOf(a.skin.rarity)
          );
        }
        case 'recent':
        default:
          return (
            new Date(b.acquired_at).getTime() - new Date(a.acquired_at).getTime()
          );
      }
    });
  }, [inventory?.items, sortBy]);

  const filteredItems = useMemo(() => {
    if (!filterRarity) return sortedItems;
    return sortedItems.filter((item) => item.skin.rarity === filterRarity);
  }, [sortedItems, filterRarity]);

  if (!mounted) {
    return null;
  }

  if (!authAPI.isAuthenticated()) {
    return null;
  }

  if (loading) {
    return (
      <div className="min-h-screen bg-background">
        <InventoryHeader />
        <div className="flex items-center justify-center h-96 text-muted-foreground">
          <Loader2 className="w-8 h-8 mr-2 animate-spin" />
          <span>Loading your inventory...</span>
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
            <Button variant="default" onClick={() => router.push('/login')}>
              Go to Login
            </Button>
          </Card>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-background">
      <InventoryHeader />

      <div className="container mx-auto px-4 py-8 text-muted-foreground bg-secondary my-10">
        <div className="mb-6 flex items-center gap-2">
          <Button variant="outline" onClick={() => router.back()}>
            <ArrowLeft className="mr-2 h-4 w-4" />
            Back
          </Button>
          <Button variant="ghost" onClick={() => router.push('/cases')}>
            Go to Cases
          </Button>
        </div>

        <div className="mb-8 rounded-lg border border-amber-500/30 bg-amber-900/10 p-4">
          <div className="mb-3 flex items-center gap-2 text-amber-300">
            <Package className="h-5 w-5" />
            <h2 className="text-lg font-semibold">Cases Ready To Open</h2>
          </div>
          {loadingPurchasedCases ? (
            <p className="text-sm text-muted-foreground">Loading purchased cases...</p>
          ) : purchasedCases.length === 0 ? (
            <p className="text-sm text-muted-foreground">
              No purchased cases yet. Buy cases from the Cases page first.
            </p>
          ) : (
            <div className="grid grid-cols-1 gap-3 sm:grid-cols-2 lg:grid-cols-3">
              {purchasedCases.map((item) => (
                <Card
                  key={item.id}
                  className="flex items-center justify-between border-gray-700 bg-gray-900/50 p-3"
                >
                  <div>
                    <p className="font-semibold text-white">{item.case.name}</p>
                    <p className="text-xs text-gray-400">
                      Bought: {new Date(item.created_at).toLocaleString()}
                    </p>
                  </div>
                  <Button
                    className="bg-blue-600 hover:bg-blue-700"
                    onClick={() => setSelectedPurchasedCase(item)}
                  >
                    Open
                  </Button>
                </Card>
              ))}
            </div>
          )}
        </div>

        <div className="mb-8 rounded-lg border border-blue-500/30 bg-blue-900/10 p-4">
          <div className="mb-3 flex items-center gap-2 text-blue-300">
            <Sparkles className="h-5 w-5" />
            <h2 className="text-lg font-semibold">AI Price Check (Mock)</h2>
          </div>
          <div className="flex flex-col gap-2 sm:flex-row">
            <input
              value={priceInput}
              onChange={(e) => setPriceInput(e.target.value)}
              placeholder="Enter skin name (e.g., AWP | Asiimov)"
              className="h-10 flex-1 rounded-md border border-gray-700 bg-gray-950 px-3 text-sm text-white"
            />
            <Button
              disabled={priceLoading || !priceInput.trim()}
              onClick={async () => {
                try {
                  setPriceLoading(true);
                  const res = await aiAPI.priceCheck(priceInput.trim());
                  setPriceResult(
                    `${res.skin_name}: $${res.suggested_usd.toFixed(2)} (${res.provider})`
                  );
                } catch (err) {
                  setPriceResult(
                    err instanceof Error ? err.message : 'Price check failed'
                  );
                } finally {
                  setPriceLoading(false);
                }
              }}
            >
              {priceLoading ? 'Checking...' : 'Check Price'}
            </Button>
          </div>
          {priceResult && <p className="mt-2 text-sm text-blue-200">{priceResult}</p>}
        </div>

        <div className="mb-6 flex items-center gap-3">
          <Button
            variant={sortBy === 'recent' ? 'default' : 'outline'}
            onClick={() => setSortBy('recent')}
          >
            Recent
          </Button>
          <Button
            variant={sortBy === 'value' ? 'default' : 'outline'}
            onClick={() => setSortBy('value')}
          >
            Value
          </Button>
          <Button
            variant={sortBy === 'rarity' ? 'default' : 'outline'}
            onClick={() => setSortBy('rarity')}
          >
            Rarity
          </Button>
          <Button variant="ghost" onClick={() => setFilterRarity(null)}>
            Clear Filter
          </Button>
        </div>

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
                      <button
                        key={rarity}
                        className="w-full flex items-center justify-between text-xs hover:opacity-80"
                        onClick={() =>
                          setFilterRarity((prev) => (prev === rarity ? null : rarity))
                        }
                      >
                        <span className="text-muted-foreground">{rarity}</span>
                        <span className="font-semibold text-foreground">{count}</span>
                      </button>
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
                <Link href="/cases">
                  <Button>Open Cases</Button>
                </Link>
              </Card>
            ) : (
              <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-3">
                {filteredItems.map((item) => (
                  <SkinCard key={item.id} item={item} />
                ))}
              </div>
            )}
          </div>
        </div>
      </div>

      {selectedPurchasedCase && (
        <CaseOpeningModal
          open={!!selectedPurchasedCase}
          onClose={() => setSelectedPurchasedCase(null)}
          caseItem={selectedPurchasedCase.case}
          purchasedCaseId={selectedPurchasedCase.id}
          onSuccess={async () => {
            await refreshAll();
            setSelectedPurchasedCase(null);
          }}
        />
      )}
    </div>
  );
}
