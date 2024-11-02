import { type AvatarProps } from '@/components/ui/avatar';
import generateRandomName from '../_utils';

export const avatars: AvatarProps[] = Array.from({ length: 10 }, () => {
  const fullNames = generateRandomName();

  return {
    src: generateRandomName(),
    name: fullNames.split(' ')[0],
  };
});
