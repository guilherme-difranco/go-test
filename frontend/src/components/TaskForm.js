import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from "react-router-dom";
import api from '../api';

const TaskForm = ({ task, refreshTasks, setEditing }) => {
    const [title, setTitle] = useState(task ? task.title : '');
    const [description, setDescription] = useState(task ? task.description : '');
    // Ajuste para usar _id em vez de id
    const [priorityId, setPriorityId] = useState(task && task.priority ? task.priority._id : '');
    const [statusId, setStatusId] = useState(task && task.status ? task.status._id : '');
    const [priorities, setPriorities] = useState([]);
    const [statuses, setStatuses] = useState([]);
    const [combosLoaded, setCombosLoaded] = useState(false);


    const { taskId } = useParams();
    const navigate = useNavigate();


    useEffect(() => {
        const fetchTaskData = async () => {
            if (taskId && combosLoaded) {
                try {
                    const { data: taskData } = await api.get(`/tasks/${taskId}`);
                    console.log(taskData)
                    setTitle(taskData.title);
                    setDescription(taskData.description);
                    setPriorityId(taskData.priority.id); // Ajuste conforme a estrutura da sua API
                    setStatusId(taskData.status.id); // Ajuste conforme a estrutura da sua API
                } catch (error) {
                    console.error('Erro ao carregar dados da tarefa', error);
                }
            }
        };

        fetchTaskData();
    }, [taskId, combosLoaded]);

    useEffect(() => {
        const fetchPrioritiesAndStatuses = async () => {
            try {
                const [priorityResponse, statusResponse] = await Promise.all([
                    api.get('/priorities'),
                    api.get('/statuses')
                ]);
                setPriorities(priorityResponse.data);
                setStatuses(statusResponse.data);
                setCombosLoaded(true);
            } catch (error) {
                console.error('Erro ao buscar priorities e statuses', error);
            }
        };
        fetchPrioritiesAndStatuses();
    }, []);

    const handleSubmit = async (e) => {
        e.preventDefault();
        const url = taskId  ? `/tasks/${taskId}` : '/task'; // Ajuste para usar _id
        const method = taskId  ? 'put' : 'post';

        const priority = priorities.find(p => p.id === priorityId);
        const status = statuses.find(s => s.id === statusId);


        if (!priority || !status) {
            console.error('Falha ao encontrar priority ou status selecionado');
            return;
        }

        const data = {
            title,
            description,
            priority: priority, // Apenas envie o _id
            status: status, // Apenas envie o _id
            creationDate: new Date(),
        };

        try {
            await api[method](url, data);
            alert('Tarefa salva com sucesso!'); // Mostra um alerta de sucesso
            setTimeout(() => navigate('/'), 3000); // Redireciona para TaskList após 3 segundos
            if (setEditing) setEditing(false);
        } catch (error) {
            console.error('Falha ao salvar a tarefa', error);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <input
                type="text"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                placeholder="Título da tarefa"
                required
            />
            <textarea
                value={description}
                onChange={(e) => setDescription(e.target.value)}
                placeholder="Descrição da tarefa"
                required
            />
            <select
                value={priorityId}
                onChange={(e) => setPriorityId(e.target.value)}
                required>
                <option value="">Selecione uma prioridade</option>
                {priorities.map((p) => (
                    <option key={p.id} value={p.id}>
                        {p.name}
                    </option>
                ))}
            </select>
            <select
                value={statusId}
                onChange={(e) => setStatusId(e.target.value)}
                required>
                <option value="">Selecione um status</option>
                {statuses.map((s) => (
                    <option key={s.id} value={s.id}>
                        {s.name}
                    </option>
                ))}
            </select>
            <button type="submit">{task ? 'Atualizar Tarefa' : 'Salvar Tarefa'}</button>
        </form>
    );
};

export default TaskForm;
