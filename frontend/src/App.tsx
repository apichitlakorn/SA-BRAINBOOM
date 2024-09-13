import { useState } from 'react';
import './App.css';
import './page/tutor/AppTutor.css';
import Header from './Header';
import Sidebar from './Sidebar';

import Tutor from './page/tutor/Tutor';
import Course from './page/course/Course';
import Calendar from './page/calendar/Calendar'; // เพิ่ม Calendar import
import CreateUser from './page/createUser/CreateUser'
import Home from './page/Dashboard/dashboard';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

function App() {
  const [openSidebarToggle, setOpenSidebarToggle] = useState<boolean>(false);

  const OpenSidebar = (): void => {
    setOpenSidebarToggle(!openSidebarToggle);
  };

  return (
    <Router>
      <div className="grid-container">
        <Header OpenSidebar={OpenSidebar} />
        <Sidebar openSidebarToggle={openSidebarToggle} OpenSidebar={OpenSidebar} />
        
        <Routes>
          <Route path='/dashboard' element = {<Home/>} />
          <Route path="/tutor" element={<Tutor />} />
          <Route path="/course" element={<Course />} />
          <Route path="/calendar" element={<Calendar />} /> 
          <Route path='/CreateUser' element = {<CreateUser/>} />
          <Route path="/" element={<Home />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
