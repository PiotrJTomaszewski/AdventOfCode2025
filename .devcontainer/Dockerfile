FROM archlinux:base-devel

COPY setup /usr/local/setup
RUN /usr/local/setup/configure.sh

RUN useradd -m -G wheel dev && echo 'dev:dev' | chpasswd
