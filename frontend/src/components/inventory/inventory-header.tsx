import Link from 'next/link';
import { Button } from '@/components/ui/button';

export default function InventoryHeader() {
  return (
    <header className="border-b border-border bg-card/30 backdrop-blur-sm sticky top-0 z-20">
      <div className="container mx-auto px-4 py-4">
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-2">
            <Link
              href="/"
              className="flex items-center gap-2 hover:opacity-80 transition"
            >
              <div className="w-8 h-8 bg-primary rounded-sm flex items-center justify-center font-bold text-primary-foreground text-sm">
                CS
              </div>
              <span className="text-xl font-bold text-foreground hidden sm:inline">
                CS:OPN
              </span>
            </Link>
          </div>
          <div className="flex items-center gap-4">
            <Link href="/inventory">
              <Button variant="ghost" size="sm">
                Inventory
              </Button>
            </Link>
            <Link href="/">
              <Button variant="ghost" size="sm">
                Cases
              </Button>
            </Link>
          </div>
        </div>
      </div>
    </header>
  );
}
