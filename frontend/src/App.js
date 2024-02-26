import React from 'react';
import { BrowserRouter as Router, Route, Link, Routes} from 'react-router-dom';
import TaskList from './components/TaskList';
import TaskForm from './components/TaskForm';

import './App.css';


function App() {
  return (
    <Router>
      <nav>
        <Link to="/">Home</Link> | <Link to="/new-task">Criar Nova Task</Link>
      </nav>
      <Routes>
        <Route path="/" element={<TaskList />} />
        <Route path="/new-task" element={<TaskForm />} />
        <Route path="/new-task/:taskId" element={<TaskForm />} /> {/* Rota para edição */}
   
      </Routes>
    </Router>
  );
}


export default App;
