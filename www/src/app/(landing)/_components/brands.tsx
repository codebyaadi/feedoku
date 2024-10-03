'use client';

import React, { useState, useEffect } from 'react';
import Image from 'next/image';
import { motion } from 'framer-motion';
import { css } from 'styled-system/css';
import {
  RedditSVG,
  StackOverflowSVG,
  MediumSVG,
  TechCrunchSVG,
  WiredSVG,
} from '../_assets/brands';

interface BrandsLogoType {
  id: string;
  src: string;
  alt:string
}

const shuffleArray = (array: BrandsLogoType[]) => {
  return array
    .map((item) => ({ ...item, sort: Math.random() }))
    .sort((a, b) => a.sort - b.sort)
    .map((item) => item);
};

const Brands = () => {
  const [shuffledBrands, setShuffledBrands] = useState(brandsLogo);

  // Shuffle the logos on every animation cycle (using a timer for this example)
  useEffect(() => {
    const interval = setInterval(() => {
      setShuffledBrands(shuffleArray(brandsLogo));
    }, 3000); // Adjust time interval as needed

    return () => clearInterval(interval); // Clean up on unmount
  }, []);

  return (
    <div
      className={css({
        w: 'full',
        px: { base: '4', lg: '10' },
        py: '4',
        mt: '20',
        display: 'flex',
        flexDir: 'column',
        justifyContent: 'space-between',
        alignItems: 'center',
        gap: '8',
      })}
    >
      <div
        className={css({
          textAlign: 'center',
        })}
      >
        Stay up to date with the world's best content platforms
      </div>
      <div
        className={css({
          display: 'grid',
          gridTemplateColumns: 'repeat(3, 1fr)',
          gridGap: '24',
          alignItems: 'center',
        })}
      >
        {shuffledBrands.map((b, idx) => (
          <motion.div
            // Use a combination of index and item id to generate a unique key on every shuffle
            key={`${b.id}-${idx}-${b.alt}-${Math.random()}`} 
            className={css({
              mx: 'auto',
            })}
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
            <Image
              src={b.src}
              alt={b.alt}
              className={css({
                h: '5',
              })}
            />
          </motion.div>
        ))}
      </div>
    </div>
  );
};

export default Brands;

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
