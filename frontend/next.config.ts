import type { NextConfig } from 'next';

const nextConfig: NextConfig = {
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'community.cloudflare.steamstatic.com',
      },
      {
        protocol: 'https',
        hostname: 'community.steamstatic.com',
      },
      {
        protocol: 'https',
        hostname: 'steamcommunity-a.akamaihd.net',
      },
      {
        protocol: 'https',
        hostname: 'community.edgecast.steamstatic.com',
      },
    ],
  },
  eslint: {
    ignoreDuringBuilds: true,
  },
};

export default nextConfig;
