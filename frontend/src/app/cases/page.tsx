'use client';

import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { authAPI, userAPI } from '@/lib/api';
import { useCases } from '@/hooks/useCase';
import { CaseCard } from '@/components/cases/CaseCard';
import { CaseOpeningModal } from '@/components/cases/CaseOpeningModal';
import { Case, User } from '@/lib/types';
import { Loader2, Package, DollarSign } from 'lucide-react';

export default function CasesPage() {
  const router = useRouter();
  const { cases, loading, error } = useCases();
  const [selectedCase, setSelectedCase] = useState<Case | null>(null);
  const [user, setUser] = useState<User | null>(null);
  const [loadingUser, setLoadingUser] = useState(true);

  useEffect(() => {
    if (!authAPI.isAuthenticated()) {
      router.push('/login');
      return;
    }

    const fetchUser = async () => {
      try {
        const userData = await userAPI.getProfile();
        console.log('👤 User data:', userData);
        setUser(userData);
      } catch (err) {
        console.error('Failed to fetch user:', err);
        router.push('/login');
      } finally {
        setLoadingUser(false);
      }
    };

    fetchUser();
  }, [router]);

  const handleCaseOpened = async () => {
    try {
      const userData = await userAPI.getProfile();
      setUser(userData);
    } catch (err) {
      console.error('Failed to refresh user:', err);
    }
  };

  if (!authAPI.isAuthenticated()) {
    return null;
  }

  if (loading || loadingUser) {
    return (
      <div className="min-h-screen bg-gray-950 flex items-center justify-center">
        <div className="text-center">
          <Loader2 className="w-12 h-12 animate-spin text-blue-500 mx-auto mb-4" />
          <p className="text-gray-400">Loading cases...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="min-h-screen bg-gray-950 flex items-center justify-center">
        <div className="text-center">
          <Package className="w-16 h-16 text-red-500 mx-auto mb-4" />
          <h2 className="text-2xl font-bold text-white mb-2">
            Error Loading Cases
          </h2>
          <p className="text-gray-400">{error}</p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-950 text-white">
      <div className="container mx-auto px-4 py-8">
        <div className="flex items-center justify-between mb-8">
          <div>
            <h1 className="text-4xl font-bold mb-2">Case Opening</h1>
            <p className="text-gray-400">Choose a case and test your luck!</p>
          </div>

          {user && user.casebucks !== undefined && (
            <div className="bg-gradient-to-r from-green-900/20 to-green-950/20 border border-green-500/30 rounded-lg p-4">
              <div className="flex items-center gap-2 text-green-400">
                <DollarSign className="w-5 h-5" />
                <div>
                  <p className="text-xs text-gray-400">Your Balance</p>
                  <p className="text-2xl font-bold">
                    {user.casebucks.toFixed(2)} CB
                  </p>
                </div>
              </div>
            </div>
          )}
        </div>

        {cases.length > 0 ? (
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
            {cases.map((caseItem) => (
              <CaseCard
                key={caseItem.id}
                caseItem={caseItem}
                onOpen={setSelectedCase}
                userBalance={user?.casebucks || 0}
              />
            ))}
          </div>
        ) : (
          <div className="text-center py-16">
            <Package className="w-20 h-20 text-gray-700 mx-auto mb-4" />
            <h3 className="text-2xl font-bold text-gray-400 mb-2">
              No Cases Available
            </h3>
            <p className="text-gray-600">Check back later for new cases!</p>
          </div>
        )}
      </div>

      {selectedCase && user && (
        <CaseOpeningModal
          open={!!selectedCase}
          onClose={() => setSelectedCase(null)}
          caseItem={selectedCase}
          onSuccess={handleCaseOpened}
        />
      )}
    </div>
  );
}
