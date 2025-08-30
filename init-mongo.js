db = db.getSiblingDB('todo-app-go');

db.tasks.insertMany([
    {
        description: "Welcome to your Todo App!"
    },
    {
        description: "Try adding your own tasks"
    }
]);

print("Database initialized successfully!");
