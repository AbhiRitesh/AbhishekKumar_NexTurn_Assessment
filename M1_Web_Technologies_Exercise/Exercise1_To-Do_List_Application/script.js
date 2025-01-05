document.addEventListener('DOMContentLoaded', () => {
    const taskInput = document.getElementById('task-input');
    const addTaskBtn = document.getElementById('add-task-btn');
    const taskList = document.getElementById('task-list');
    const pendingCount = document.getElementById('pending-count');
  
    let tasks = JSON.parse(localStorage.getItem('tasks')) || [];
  
    const renderTasks = () => {
      taskList.innerHTML = '';
      tasks.forEach((task, index) => {
        const li = document.createElement('li');
        li.className = `task ${task.completed ? 'completed' : ''}`;
        li.setAttribute('draggable', true);
        li.dataset.index = index;
        li.innerHTML = `
          <span>${task.name}</span>
          <div class="task-actions">
            <button class="complete-btn" data-index="${index}">✔</button>
            <button class="edit-btn" data-index="${index}">✏</button>
            <button class="delete-btn" data-index="${index}">✖</button>
          </div>
        `;
        taskList.appendChild(li);
  
        // Add drag event listeners
        li.addEventListener('dragstart', handleDragStart);
        li.addEventListener('dragover', handleDragOver);
        li.addEventListener('drop', handleDrop);
        li.addEventListener('dragend', handleDragEnd);
      });
      updatePendingCount();
    };
  
    const updatePendingCount = () => {
      const pendingTasks = tasks.filter(task => !task.completed).length;
      pendingCount.textContent = pendingTasks;
    };
  
    const saveTasks = () => {
      localStorage.setItem('tasks', JSON.stringify(tasks));
    };
  
    addTaskBtn.addEventListener('click', () => {
      const taskName = taskInput.value.trim();
      if (taskName) {
        tasks.push({ name: taskName, completed: false });
        taskInput.value = '';
        saveTasks();
        renderTasks();
      }
    });
  
    taskList.addEventListener('click', (e) => {
      const index = e.target.dataset.index;
      if (e.target.classList.contains('complete-btn')) {
        tasks[index].completed = !tasks[index].completed;
      } else if (e.target.classList.contains('edit-btn')) {
        const newName = prompt('Edit task:', tasks[index].name);
        if (newName !== null) {
          tasks[index].name = newName.trim();
        }
      } else if (e.target.classList.contains('delete-btn')) {
        tasks.splice(index, 1);
      }
      saveTasks();
      renderTasks();
    });
  
    // Drag-and-drop event handlers
    let draggedElement;
  
    const handleDragStart = (e) => {
      draggedElement = e.target;
      draggedElement.classList.add('dragging');
    };
  
    const handleDragOver = (e) => {
      e.preventDefault();
      const target = e.target.closest('.task');
      if (target && target !== draggedElement) {
        target.classList.add('drop-target');
      }
    };
  
    const handleDrop = (e) => {
      e.preventDefault();
      const target = e.target.closest('.task');
      if (target && target !== draggedElement) {
        const draggedIndex = parseInt(draggedElement.dataset.index);
        const targetIndex = parseInt(target.dataset.index);
  
        // Swap tasks in the array
        const temp = tasks[draggedIndex];
        tasks[draggedIndex] = tasks[targetIndex];
        tasks[targetIndex] = temp;
  
        saveTasks();
        renderTasks();
      }
    };
  
    const handleDragEnd = (e) => {
      draggedElement.classList.remove('dragging');
      document.querySelectorAll('.task').forEach((task) => {
        task.classList.remove('drop-target');
      });
    };
  
    renderTasks();
  });
  