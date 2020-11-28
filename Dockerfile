FROM alpine
ADD microdemo /microdemo
ENTRYPOINT [ "/microdemo" ]
