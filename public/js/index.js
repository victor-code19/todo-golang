const formInput = document.getElementById('form_input')
const inputField = document.getElementById('input_field')
const addButton = document.getElementById('add_button')
const clearAllBtn = document.getElementById('clear_all_btn')
const todoList = document.getElementById('todo_list')
const info = document.getElementsByClassName('info')


inputField.onkeyup = () => {
    let userData = inputField.value.trim()
    if (userData) {
        addButton.classList.add('active')
    } else {
        addButton.classList.remove('active')
    }
}

formInput.addEventListener('submit', async (e) => {
    e.preventDefault()
    const taskData = inputField.value

    const response = await fetch("/api/task", {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            description: taskData
        })
    })

    if (response.status === 201) {
        const data = await response.json()
        
        const taskElement = document.createElement('li')
        const button = document.createElement('button')
        const trashIcon = document.createElement('i')
        
        trashIcon.classList.add('fa', 'fa-trash')

        button.appendChild(trashIcon)
        button.setAttribute('id', data.id)
        button.setAttribute('onclick', 'deleteItem(this.id)')

        taskElement.appendChild(document.createTextNode(data.description))
        taskElement.appendChild(button)

        todoList.appendChild(taskElement)
        inputField.value = ""
        addButton.classList.remove('active')

        getTasksAmountInfo()
        console.log(data.id)
    }
    else {
        info[0].textContent = "Unable to add task."
    }         
})

clearAllBtn.addEventListener('click', async () => {
    const response = await fetch("/api/tasks", {
        method: 'DELETE'
    })

    if (response.status === 200) {
        while (todoList.firstChild) {
            todoList.removeChild(todoList.firstChild)
        }

        getTasksAmountInfo()
    }
    else {
        info[0].textContent = "Unable to delete tasks."
    }

})

async function deleteItem(id) {
    const response = await fetch(`/api/task/${id}`, {
         method: 'DELETE'
    })

    if (response.status === 200) {
        const task = document.getElementById(id)
        task.parentElement.remove()

        getTasksAmountInfo()
    } else {
        info[0].textContent = "Unable to delete task."
    }
}

function getTasksAmountInfo() {
    const tasks = todoList.getElementsByTagName("li").length
    if (!tasks) {
        info[0].textContent = "No tasks available."
    } else {
        info[0].textContent = `You have ${tasks} pending tasks`
    }
}