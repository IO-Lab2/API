# API
API component for the project

## How to work with devcontainer

### Prerequisites
- Docker
- Visual Studio Code
- Visual Studio Code Remote - Containers extension

### Steps
1. Clone the repository and open it in Visual Studio Code
2. Open the command palette (Ctrl+Shift+P) and run `Dev Container: Reopen in Container`
3. Wait for the container to build and open the project in the container

### Awesome docs:
https://api.epickaporownywarkabazwiedzyuczelni.rocks/docs#/

### Workflow:
Po zmianach w repo automatycznie uruchamia się workflow z testami (Go) i jeżeli zakończy się on pomyślnie to automatycznie jest uruchamiany workflow dockerImage tworzący obraz API, który potem jest wrzucany na serwer.

Jeżeli potrzebujesz, żeby mimo nie przechodzenia testów wrzucić nową wersje API na serwer to w zakładce 'Actions' wybierasz 'dockerImage' z lewej strony i z prawej strony masz przycisk 'Run workflow', klikasz go, nic nie zmieniasz, znowu klikasz 'Run workflow' i wszystko się samo robi.

Watchtower na serwerze sprawdza czy jest nowa wersja obrazu API co dziesięć minut, więc po co najwyrzej 10 minutach od zakończenia workflowu dockerImage nowa wersja powinna się pojawić na werwerze.
