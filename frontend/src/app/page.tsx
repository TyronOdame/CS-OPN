import { Button } from '@/components/ui/button';
import { Card } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';

export default function HomePage() {
  return (
    <div className="min-h-screen bg-[#222222]">
      {/* Navigation */}
      <nav className="border-b border-border bg-card/50 backdrop-blur-sm">
        <div className="container mx-auto px-4 py-4">
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-2">
              <div className="w-8 h-8 bg-primary rounded-sm flex items-center justify-center font-bold text-primary-foreground">
                CS
              </div>
              <span className="text-xl font-bold text-foreground">CS:OPN</span>
            </div>
            <div className="flex items-center gap-6">
              <a
                href="#"
                className="text-muted-foreground hover:text-foreground transition-colors"
              >
                Cases
              </a>
              <Button variant="outline" size="sm">
                Sign In
              </Button>
            </div>
          </div>
        </div>
      </nav>

      {/* Hero Section */}
      <section className="relative py-20 px-4 overflow-hidden">
        <div className="absolute inset-0 bg-gradient-to-br from-primary/10 via-[#222222] to-[#222222]"></div>
        <div className="container mx-auto text-center relative z-10">
          <Badge className="mb-6 bg-primary/20 text-primary border-primary/30">
            New: Operation Bravo Cases Available
          </Badge>
          <h1 className="text-5xl md:text-7xl font-bold mb-6 text-balance">
            Open CS2 Cases
            <span className="block text-primary">Risk Free</span>
          </h1>
          <p className="text-xl text-muted-foreground mb-8 max-w-2xl mx-auto text-pretty">
            Experience the thrill of case opening without spending real money.
            Collect rare skins, build your inventory, and trade with friends.
          </p>
          <div className="flex justify-center">
            <Button size="lg" className="cs2-glow text-lg px-8 py-6">
              Start Opening Cases
            </Button>
          </div>
        </div>
      </section>

      {/* Featured Cases */}
      <section className="py-16 px-4">
        <div className="container mx-auto">
          <h2 className="text-3xl font-bold text-center mb-12">
            Featured Cases
          </h2>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            {[
              {
                name: 'Operation Bravo Case',
                rarity: 'Special',
                color: 'text-primary',
              },
              {
                name: 'Chroma 3 Case',
                rarity: 'Rare',
                color: 'text-chart-2',
              },
              {
                name: 'Fracture Case',
                rarity: 'Common',
                color: 'text-chart-4',
              },
            ].map((caseItem, index) => (
              <Card
                key={index}
                className="p-6 hover:bg-card/80 transition-all duration-300 cs2-float group cursor-pointer"
              >
                <div className="text-center">
                  <div className="mb-4 relative">
                    <div className="w-32 h-32 mx-auto bg-muted/20 rounded-lg flex items-center justify-center group-hover:scale-110 transition-transform duration-300">
                      <span className="text-muted-foreground text-sm">
                        Case Image
                      </span>
                    </div>
                    <div className="absolute inset-0 bg-gradient-to-t from-primary/20 to-transparent rounded-lg opacity-0 group-hover:opacity-100 transition-opacity"></div>
                  </div>
                  <h3 className="text-lg font-semibold mb-2">
                    {caseItem.name}
                  </h3>
                  <Badge
                    variant="outline"
                    className={`mb-4 ${caseItem.color} border-current`}
                  >
                    {caseItem.rarity}
                  </Badge>
                  <Button className="w-full">Open Case</Button>
                </div>
              </Card>
            ))}
          </div>
        </div>
      </section>

      {/* Stats Section */}
      <section className="py-16 px-4 bg-card/30">
        <div className="container mx-auto">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8 text-center">
            {[
              { label: 'Cases Opened', value: '2.4M+' },
              { label: 'Rare Skins Found', value: '156K+' },
              { label: 'Total Value', value: '$12.8M+' },
            ].map((stat, index) => (
              <div key={index}>
                <div className="text-3xl md:text-4xl font-bold text-primary mb-2">
                  {stat.value}
                </div>
                <div className="text-muted-foreground">{stat.label}</div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* How It Works */}
      <section className="py-16 px-4">
        <div className="container mx-auto">
          <h2 className="text-3xl font-bold text-center mb-12">How It Works</h2>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            {[
              {
                step: '1',
                title: 'Choose Your Case',
                description:
                  'Select from dozens of CS2 cases with authentic drop rates and skin collections.',
              },
              {
                step: '2',
                title: 'Open & Discover',
                description:
                  'Experience the authentic case opening animation and discover your new skin.',
              },
              {
                step: '3',
                title: 'Build Collection',
                description:
                  'Add skins to your inventory, trade with friends, or sell on the marketplace.',
              },
            ].map((item, index) => (
              <Card key={index} className="p-6 text-center">
                <div className="w-12 h-12 bg-primary text-primary-foreground rounded-full flex items-center justify-center text-xl font-bold mx-auto mb-4">
                  {item.step}
                </div>
                <h3 className="text-xl font-semibold mb-3">{item.title}</h3>
                <p className="text-muted-foreground text-pretty">
                  {item.description}
                </p>
              </Card>
            ))}
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-20 px-4 bg-gradient-to-r from-primary/10 to-primary/5">
        <div className="container mx-auto text-center">
          <h2 className="text-4xl font-bold mb-6">Ready to Start Opening?</h2>
          <p className="text-xl text-muted-foreground mb-8 max-w-2xl mx-auto">
            Join thousands of players experiencing the thrill of CS2 case
            opening. No risk, all the excitement.
          </p>
          <Button size="lg" className="text-lg px-8 py-6 cs2-glow">
            Open Your First Case
          </Button>
        </div>
      </section>

      {/* Footer */}
      <footer className="border-t border-border py-12 px-4">
        <div className="container mx-auto">
          <div className="grid grid-cols-2 md:grid-cols-4 gap-8">
            <div>
              <h4 className="font-semibold mb-4">Cases</h4>
              <ul className="space-y-2 text-sm text-muted-foreground">
                <li>
                  <a
                    href="#"
                    className="hover:text-foreground transition-colors"
                  >
                    All Cases
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-foreground transition-colors"
                  >
                    New Releases
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-foreground transition-colors"
                  >
                    Popular
                  </a>
                </li>
              </ul>
            </div>
            <div>
              <h4 className="font-semibold mb-4">Features</h4>
              <ul className="space-y-2 text-sm text-muted-foreground">
                <li>
                  <a
                    href="#"
                    className="hover:text-foreground transition-colors"
                  >
                    Inventory
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-foreground transition-colors"
                  >
                    Trading
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-foreground transition-colors"
                  >
                    Marketplace
                  </a>
                </li>
              </ul>
            </div>
            <div>
              <h4 className="font-semibold mb-4">Support</h4>
              <ul className="space-y-2 text-sm text-muted-foreground">
                <li>
                  <a
                    href="#"
                    className="hover:text-foreground transition-colors"
                  >
                    Help Center
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-foreground transition-colors"
                  >
                    Contact
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-foreground transition-colors"
                  >
                    FAQ
                  </a>
                </li>
              </ul>
            </div>
            <div>
              <h4 className="font-semibold mb-4">Legal</h4>
              <ul className="space-y-2 text-sm text-muted-foreground">
                <li>
                  <a
                    href="#"
                    className="hover:text-foreground transition-colors"
                  >
                    Privacy
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-foreground transition-colors"
                  >
                    Terms
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-foreground transition-colors"
                  >
                    Disclaimer
                  </a>
                </li>
              </ul>
            </div>
          </div>
          <div className="border-t border-border mt-8 pt-8 text-center text-sm text-muted-foreground">
            <p>
              © 2025 CS2 Case Simulator. Not affiliated with Valve Corporation.
            </p>
          </div>
        </div>
      </footer>
    </div>
  );
}
