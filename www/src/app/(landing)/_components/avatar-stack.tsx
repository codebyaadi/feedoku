import { Avatar, type AvatarProps } from '@/components/ui/avatar';
import React from 'react';
import { css } from 'styled-system/css';

interface AvatarStackProps {
  avatars: AvatarProps[];
  maxVisible?: number;
}

const AvatarStack: React.FC<AvatarStackProps> = ({
  avatars,
  maxVisible = 5,
}) => {
  const visibleAvatars = avatars.slice(0, maxVisible);
  const remainingCount = avatars.length - maxVisible;

  return (
    <div
      className={css({
        display: 'flex',
        position: 'relative',
      })}
    >
      {visibleAvatars.map((avatar, index) => (
        <Avatar
          key={index}
          src={avatar.src}
          name={avatar.name}
          className={css({
            ml: '-4',
          })}
        />
      ))}
      {remainingCount > 0 && (
        <Avatar
          name={`+ ${remainingCount}`}
          className={css({
            ml: '-4',
          })}
        />
      )}
    </div>
  );
};

export default AvatarStack;
