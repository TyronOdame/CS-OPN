import Image from 'next/image';

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
}

interface SkinCardProps {
  item: InventoryItem;
}

const FALLBACK_IMAGE_SRC = '/file.svg';

const RARITY_COLORS: Record<
  string,
  { bg: string; border: string; text: string }
> = {
  Covert: {
    bg: 'from-red-950 to-red-900',
    border: 'border-red-600',
    text: 'text-red-400',
  },
  Classified: {
    bg: 'from-purple-950 to-purple-900',
    border: 'border-purple-600',
    text: 'text-purple-400',
  },
  Restricted: {
    bg: 'from-blue-950 to-blue-900',
    border: 'border-blue-600',
    text: 'text-blue-400',
  },
  'Mil-Spec': {
    bg: 'from-cyan-950 to-cyan-900',
    border: 'border-cyan-600',
    text: 'text-cyan-400',
  },
  'Industrial Grade': {
    bg: 'from-amber-950 to-amber-900',
    border: 'border-amber-600',
    text: 'text-amber-400',
  },
  'Consumer Grade': {
    bg: 'from-gray-950 to-gray-900',
    border: 'border-gray-600',
    text: 'text-gray-400',
  },
};

export default function SkinCard({ item }: SkinCardProps) {
  const colors =
    RARITY_COLORS[item.skin.rarity] || RARITY_COLORS['Consumer Grade'];

  return (
    <div
      className={`group relative cursor-pointer overflow-hidden border-2 ${colors.border} bg-gradient-to-br ${colors.bg} aspect-square hover:shadow-lg transition-all duration-300 hover:-translate-y-1`}
    >
      {/* Image Container */}
      {item.skin.image_url ? (
        <Image
          src={item.skin.image_url || FALLBACK_IMAGE_SRC}
          alt={item.skin.name}
          fill
          unoptimized
          onError={(event) => {
            const img = event.currentTarget as HTMLImageElement;
            if (img.dataset.fallbackApplied) return;
            img.dataset.fallbackApplied = 'true';
            img.src = FALLBACK_IMAGE_SRC;
          }}
          className="object-cover group-hover:scale-110 transition-transform duration-300"
        />
      ) : (
        <div className="w-full h-full flex items-center justify-center bg-card/30">
          <p className="text-muted-foreground text-xs text-center px-2">
            {item.skin.name}
          </p>
        </div>
      )}

      {/* Overlay on Hover */}
      <div className="absolute inset-0 bg-gradient-to-t from-black/80 via-black/40 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300 flex flex-col justify-end p-2">
        <p className="text-white text-xs font-semibold truncate">
          {item.skin.name}
        </p>
        <p className="text-primary text-xs font-bold">
          ${item.value.toFixed(2)}
        </p>
        <p className="text-muted-foreground text-xs">{item.condition}</p>
      </div>

      {/* Rarity indicator corner */}
      <div className="absolute top-1 right-1 opacity-0 group-hover:opacity-100 transition-opacity duration-300">
        <span
          className={`inline-block px-1.5 py-0.5 text-xs font-bold rounded ${colors.text} bg-black/50`}
        >
          {item.skin.rarity.charAt(0)}
        </span>
      </div>
    </div>
  );
}
