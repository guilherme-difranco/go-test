import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import api from '../api';
import DeleteButton from './DeleteButton';

const TaskList = () => {
  const [tasks, setTasks] = useState([]);
  const navigate = useNavigate(); // Adicionado para navegação

  const refreshTasks = async () => {
    try {
      const { data } = await api.get('/tasks');
      setTasks(data);
    } catch (error) {
      console.error('Erro ao buscar tarefas', error);
    }
  };

  useEffect(() => {
    refreshTasks();
  }, []);

  const handleEdit = (taskId) => {
    navigate(`/new-task/${taskId}`); // Redireciona para a tela de edição com o ID da tarefa
  };

  return (
    <table>
      <thead>
        <tr>
          <th>Título</th>
          <th>Descrição</th>
          <th>Prioridade</th>
          <th>Status</th>
          <th>Data de Criação</th>
          <th>Ações</th>
        </tr>
      </thead>
      <tbody>
        {tasks.map(task => (
          <tr key={task.id}>
            <td>{task.title}</td>
            <td>{task.description}</td>
            <td>{task.priority.name}</td>
            <td>{task.status.name}</td>
            <td>{new Date(task.creationDate).toLocaleDateString()}</td>
            <td>
              <button onClick={() => handleEdit(task.id)}>Editar</button>
              <DeleteButton taskId={task.id} refreshTasks={refreshTasks} />
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default TaskList;
