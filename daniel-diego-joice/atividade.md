## Desafio Técnico – API de Agendamento de Consultas Médicas

---

### **Contexto do Cliente**

Uma clínica médica chamada **SaúdeMais** deseja desenvolver uma aplicação web onde **pacientes possam se cadastrar e agendar suas próprias consultas médicas**. O sistema deve garantir que **somente o próprio paciente tenha acesso às suas consultas**, com autenticação segura via JWT. Por enquanto, não será necessário cadastrar médicos nem especialidades.

O sistema deve permitir:

- Cadastro e login de pacientes.
- Agendamento, consulta e cancelamento de consultas.
- Validações para impedir agendamentos no passado.

---

### **Requisitos Funcionais**

### 1. **Cadastro de Paciente**

- `POST /register` com:
  - `name`: string, obrigatório
  - `email`: string, obrigatório e único
  - `password`: string, obrigatório
- Resposta: mensagem de sucesso ou erro padronizado.

### 2. **Login**

- `POST /login` com:
  - `email`
  - `password`
- Resposta:

  ```json
  json
  {
    "token": "JWT_AQUI"
  }

  ```

### 3. **Agendar Consulta**

- `POST /appointments`
- Requer JWT no header `Authorization`.
- Body:

  ```json
  json
  {
    "datetime": "2025-06-20T14:30:00"
  }

  ```

- Não permitir:
  - Agendar no passado.
  - Duas consultas no mesmo horário para o mesmo paciente.

### 4. **Listar Consultas**

- `GET /appointments`
- Requer JWT.
- Retorna todas as consultas agendadas do paciente autenticado.

### 5. **Cancelar Consulta**

- `DELETE /appointments/:id`
- Requer JWT.
- Só pode cancelar consultas do próprio usuário.

### **Mensagens de Erro Padrão**

```json
json
{
  "message": "Paciente já cadastrado",
  "code": 409
}

```

```json
json
{
  "message": "Consulta no passado não é permitida",
  "code": 400
}

```

```json
json
{
  "message": "Consulta não encontrada ou acesso não autorizado",
  "code": 403
}

```

```json
json
{
  "message": "Token inválido ou expirado",
  "code": 401
}

```

---

### **Critérios de Aceitação**

- Criar paciente e autenticar → Receber JWT
- Agendar consulta → Sucesso
- Agendar duas no mesmo horário → Erro
- Agendar no passado → Erro
- Listar consultas → Retornar somente as do paciente logado
- Cancelar consulta → Funciona apenas com ID válido do próprio usuário
