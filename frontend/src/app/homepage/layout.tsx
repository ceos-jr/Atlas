import SideBar from '@/app/ui/sidebar';
 
export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex h-screen flex-col bg-white md:flex-row md:overflow-hidden">
      <div className="w-full flex-none md:w-64">
        <SideBar />
      </div>
      <div className="flex-grow p-6 md:overflow-y-auto md:p-12">{children}</div>
    </div>
  );
}