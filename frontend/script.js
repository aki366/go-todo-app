document.addEventListener('DOMContentLoaded', function() {
  fetchTodos();

  async function fetchTodos() {
      const response = await fetch('http://localhost:8080/todos');
      const todos = await response.json();
      const todoList = document.getElementById('todoList');
      todoList.innerHTML = '';
      todos.forEach(todo => {
          const li = document.createElement('li');
          li.textContent = todo.description;
          const deleteButton = document.createElement('button');
          deleteButton.textContent = 'Delete';
          deleteButton.onclick = () => deleteTodo(todo.id);
          li.appendChild(deleteButton);
          todoList.appendChild(li);
      });
  }

  window.addTodo = async function() {
      const input = document.getElementById('newTodo');
      const description = input.value;
      const response = await fetch('http://localhost:8080/todos', {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json'
          },
          body: JSON.stringify({ description: description, completed: false })
      });
      if (response.ok) {
          fetchTodos();
          input.value = ''; // clear input after adding
      }
  };

  window.deleteTodo = async function(id) {
      const response = await fetch(`http://localhost:8080/todos/${id}`, {
          method: 'DELETE'
      });
      if (response.ok) {
          fetchTodos();
      }
  };
});
