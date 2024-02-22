import React from 'react';

const SideBar = () => {
  return (
    
    <div
        className="relative pattern"
    >
        <nav
        className="z-20 flex shrink-0 grow-0 justify-around gap-4 border-t border-gray-200 bg-white/50 p-8 shadow-lg m-4 backdrop-blur-lg bg-neutral-200 fixed top-2/4 -translate-y-2/4 left-6 min-h-[auto] min-w-[64px] flex-col rounded-lg border"
        >
        <a
            href="#logo"
            className="flex aspect-square min-h-[32px] w-16 flex-col items-center justify-center gap-1 rounded-md p-8 bg-indigo-50 text- dark:bg-sky-900 dark:text-sky-50"
        >
            <small className="text-center text-xs font-medium"> </small>

        </a>
    
        <a
            href="#homepage"
            className="flex aspect-square min-h-[32px] w-16 flex-col items-center justify-center gap-1 rounded-md p-8 text-gray-700 hover:bg-blue-300 shadow-md"
        >
    
            <small className="text-center text-xs font-medium"> Home </small>
        </a>
    
        <a
            href="#meuperfil"
            className="flex aspect-square min-h-[32px] w-16 flex-col items-center justify-center gap-1 rounded-md p-8 text-gray-700 hover:bg-blue-300 shadow-md"
        >
           
            <small className="text-center text-xs font-medium"> Meu Perfil </small>
        </a>
    
        
    
        <a
            href="#equipes"
            className="flex aspect-square min-h-[32px] w-16 flex-col items-center justify-center gap-1 rounded-md p-8 text-gray-700 hover:bg-blue-300 shadow-md"
        >
            <small className="text-xs font-medium">Equipes</small>
        </a>

        <a
            href="#meuspontos"
            className="flex aspect-square min-h-[32px] w-16 flex-col items-center justify-center gap-1 rounded-md p-8 text-gray-700 hover:bg-blue-300 shadow-md"
        >
           
            <small className="text-center text-xs font-medium"> Meus pontos </small>
        </a>
        <a
            href="#config"
            className="flex aspect-square min-h-[32px] w-16 flex-col items-center justify-center gap-1 rounded-md p-8 text-gray-700 hover:bg-blue-300 shadow-md"
        >
           
            <small className="text-center text-xs font-medium"> Configurações </small>
        </a>
        <a
            href="#logout"
            className="flex aspect-square min-h-[32px] w-16 flex-col items-center justify-center gap-1 rounded-md p-8 text-gray-700 hover:bg-blue-300 shadow-md"
        >
           
            <small className="text-center text-xs font-medium"> Sair </small>
        </a>
        </nav>
    </div>
  );
};

export default SideBar;
