import { grid } from 'styled-system/patterns';

export default function AuthLayout() {
  return (
    <main
      className={grid({
        gridTemplateColumns: { base: 1, lg: 2 },
        maxW: '7xl',
        mx: { base: '4', lg: 'auto' },
      })}
    >
      <div>Grid 1</div>
      <div>Grid 2</div>
    </main>
  );
}
