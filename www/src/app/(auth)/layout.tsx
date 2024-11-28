export default function AuthLayout({
    children,
  }: {
    children: React.ReactNode;
  }) {
    return (
      <div className='relative pt-16 font-outfit mx-auto'>
        {children}
      </div>
    );
  }