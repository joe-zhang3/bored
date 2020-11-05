FROM python:3.9-slim as base

RUN pip install pipenv

COPY Pipfile .
COPY Pipfile.lock .

RUN PIPENV_VENV_IN_PROJECT=1 pipenv install --deploy

FROM base AS runtime

COPY --from=base /.venv /.venv

ENV PATH="/.venv/bin:$PATH"

RUN useradd --create-home appuser
WORKDIR /home/appuser
USER appuser

COPY . .

# Run the application
CMD ["python", "./main.py"]
