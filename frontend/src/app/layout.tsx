import "./globals.css";

export const metadata = {
  title: "Atlas - Ceos Jr.",
  description: "Plataforma de gerenciamento de equipes e demandas.",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className="text-3xl bg-slate-200-800">{children}</body>
    </html>
  );
}
