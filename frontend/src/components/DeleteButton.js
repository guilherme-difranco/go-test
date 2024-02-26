import React from 'react';
import api from '../api';

const DeleteButton = ({ taskId, refreshTasks }) => {
  const handleDelete = async () => {
    try {
      await api.delete(`/tasks/${taskId}`);
      refreshTasks(); // Função para atualizar a lista de tarefas
    } catch (error) {
      console.error('Falha ao deletar a tarefa', error);
    }
  };

  return (
    <button onClick={handleDelete}>Deletar</button>
  );
};

export default DeleteButton;
