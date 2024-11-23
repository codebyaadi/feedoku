'use client';

import { useEffect, useState } from 'react';
import {
  MediumSVG,
  RedditSVG,
  StackOverflowSVG,
  TechCrunchSVG,
  WiredSVG,
} from '../_assets/brands';
import { motion } from 'framer-motion';
import Image from 'next/image';

interface BrandsLogoType {
  id: string;
  src: string;
  alt: string;
}

const shuffleArray = (array: BrandsLogoType[]) => {
  return array
    .map((item) => ({ ...item, sort: Math.random() }))
    .sort((a, b) => a.sort - b.sort)
    .map((item) => item);
};

export const Brands = () => {
  const [shuffledBrands, setShuffledBrands] = useState(brandsLogo);

  // Shuffle the logos on every animation cycle (using a timer for this example)
  useEffect(() => {
    const interval = setInterval(() => {
      setShuffledBrands(shuffleArray(brandsLogo));
    }, 3000); // Adjust time interval as needed

    return () => clearInterval(interval); // Clean up on unmount
  }, []);

  return (
    <div className='mt-20 flex w-full flex-col items-center justify-between gap-12 px-4 py-4 lg:px-10'>
      <div className='text-center'>
        Stay up to date with the world's best content platforms
      </div>
      <div className='grid grid-cols-3 items-center gap-24'>
        {shuffledBrands.map((b, idx) => (
          <motion.div
            key={`${b.id}-${idx}-${b.alt}-${Math.random()}`}
            className='mx-auto'
            layout
            initial={{ opacity: 0, scale: 0.5, filter: 'blur(10px)' }}
            exit={{ opacity: 0, scale: 0.5, filter: 'blur(10px)' }}
            animate={{
              opacity: 1,
              scale: 1,
              filter: 'blur(0px)',
            }}
            transition={{
              duration: 1.5,
              ease: 'easeInOut',
            }}
          >
            <Image src={b.src} alt={b.alt} className='h-5' />
          </motion.div>
        ))}
      </div>
    </div>
  );
};

const brandsLogo: BrandsLogoType[] = [
  {
    id: '1',
    src: RedditSVG,
    alt: 'reddit-logo',
  },
  {
    id: '2',
    src: StackOverflowSVG,
    alt: 'stackoverflow-logo',
  },
  {
    id: '3',
    src: MediumSVG,
    alt: 'medium-logo',
  },
  {
    id: '4',
    src: TechCrunchSVG,
    alt: 'techcrunch-logo',
  },
  {
    id: '5',
    src: WiredSVG,
    alt: 'wired-logo',
  },
  {
    id: '6',
    src: RedditSVG,
    alt: 'reddit1-svg',
  },
];
