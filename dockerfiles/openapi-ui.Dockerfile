FROM swaggerapi/swagger-ui:v5.32.8@sha256:1ece6f19ca5709fb5fbab9bf43adee2916144fb21e82b7bafba568124e8c8930

COPY docs/openapi.yaml /docs/openapi.yaml

ENV SWAGGER_JSON=/docs/openapi.yaml
