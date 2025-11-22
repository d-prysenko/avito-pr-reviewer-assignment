# syntax=docker/dockerfile:1

FROM gomicro/goose

ADD migrations/. /migrations/

ENTRYPOINT ["goose", "postgres", "up"]
