package generated

//go:generate rm -rf models restapi client
//go:generate swagger generate server -f ../../../../swagger.yml --strict-responders --strict-additional-properties --keep-spec-order --principal github.com/Meat-Hook/point-bank/internal/modules/user/internal/app.Session --exclude-main
//go:generate swagger generate client -f ../../../../swagger.yml
