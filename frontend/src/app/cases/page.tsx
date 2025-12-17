'use client';

import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { authAPI, userAPI } from '@/lib/api';
import { useCases } from '@/hooks/useCase';
import { CaseCard } from '@/components/cases/CaseCard';
import { CaseOpeningModal } from '@/components/cases/CaseOpeningModal';
import { Case, User } from '@/lib/types';
import { Loader2, Package, DollarSign } from 'lucide-react';
