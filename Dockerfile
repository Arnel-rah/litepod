FROM codercom/code-server:latest

USER root
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - && \
  apt-get install -y nodejs && \
  npm install -g pnpm

USER coder
ENV PATH="/home/coder/.local/bin:$PATH"

WORKDIR /home/coder/project
