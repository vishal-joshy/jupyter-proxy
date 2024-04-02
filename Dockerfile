FROM jupyter/base-notebook

ADD jupyterhub_config.py /srv/jupyterhub/jupyterhub_config.py

EXPOSE 8000

CMD ["jupyterhub", "--config", "/srv/jupyterhub/jupyterhub_config.py"]
