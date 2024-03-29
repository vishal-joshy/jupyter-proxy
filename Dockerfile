FROM jupyter/base-notebook

RUN pip install jupyterhub

COPY jupyterhub_config.py /srv/jupyterhub/

EXPOSE 8000

CMD ["jupyterhub", "--config", "/srv/jupyterhub/jupyterhub_config.py"]
