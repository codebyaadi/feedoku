import { BrandsLogoType, CustomSVGProps } from '../_types';
import MediumSVG from './medium-logo';
import RedditSVG from './reddit-logo';
import StackOverflowSVG from './stackoverflow-logo';
import TechCrunchSVG from './techcrunch-logo';
import WiredSVG from './wired-logo';

export const brandsLogo: BrandsLogoType<CustomSVGProps>[] = [
  {
    id: '1',
    fc: RedditSVG,
    alt: 'reddit-logo',
  },
  {
    id: '2',
    fc: StackOverflowSVG,
    alt: 'stackoverflow-logo',
  },
  {
    id: '3',
    fc: MediumSVG,
    alt: 'medium-logo',
  },
  {
    id: '4',
    fc: TechCrunchSVG,
    alt: 'techcrunch-logo',
  },
  {
    id: '5',
    fc: WiredSVG,
    alt: 'wired-logo',
  },
  {
    id: '6',
    fc: RedditSVG,
    alt: 'reddit1-svg',
  },
];
