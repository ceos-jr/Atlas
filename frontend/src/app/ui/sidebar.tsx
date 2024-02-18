import React from 'react';

const SideBar = () => {
  return (
    <div className="flex flex-col h-screen bg-gray-800 text-white">
      <div className="p-4 hover:bg-gray-700 cursor-pointer">
        <i className="fas fa-user mr-2"></i> User
      </div>
      <div className="p-4 hover:bg-gray-700 cursor-pointer">
        <i className="fas fa-home mr-2"></i> Home
      </div>
      {/* Add more items as needed */}
    </div>
  );
};

export default SideBar;
