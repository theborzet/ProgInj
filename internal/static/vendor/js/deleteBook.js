document.addEventListener('DOMContentLoaded', () => {
    const deleteButtons = document.querySelectorAll('.delete-btn');

    deleteButtons.forEach(button => {
        button.addEventListener('click', async () => {
            const bookId = button.getAttribute('data-book-id');
            try {
                // Отправляем DELETE-запрос на сервер
                const response = await fetch(`/books/${bookId}`, {
                    method: 'DELETE'
                });
                if (response.ok) {
                    // Если запрос выполнен успешно, перезагружаем страницу
                    window.location.reload();
                } else {
                    window.location.reload();
                }
            } catch (error) {
                console.error('Ошибка удаления книги:', error.message);
            }
        });
    });
});
