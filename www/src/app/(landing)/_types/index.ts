export interface CustomSVGProps extends React.SVGProps<SVGSVGElement> {
  fillcolor: string;
}

export interface BrandsLogoType<T extends React.SVGProps<SVGSVGElement>> {
  id: string;
  fc: React.ComponentType<T>;
  alt: string;
}
