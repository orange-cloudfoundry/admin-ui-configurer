# Admin-ui-configurer

Wrap [admin-ui](https://github.com/cloudfoundry-incubator/admin-ui) to make it runnable simply as a cloud foundry app.

## Deploy

1. Download latest `admin-ui-configurer-bundled.zip` in [release page](https://github.com/orange-cloudfoundry/admin-ui-configurer/releases)
2. Create a mysql service on your cloud foundry where you want to deploy the app (e.g.: `cf cs p-mysql 100mb admin-ui-mysql`)
3. Create an uaa client for your admin-ui, e.g.:
```
uaac client add admin_ui_client \
     --authorities clients.write,cloud_controller.admin,cloud_controller.read,cloud_controller.write,doppler.firehose,openid,scim.read,scim.write,sps.write \
     --authorized_grant_types authorization_code,client_credentials,refresh_token \
     --autoapprove true \
     --scope admin_ui.admin,admin_ui.user,openid \
     -s myclientsecret
```
4. Extract the previous downloaded zip file and set `manifest.yml` and `config.json` as you want
5. Push your config as a service with this command `cf cups admin-ui-config -p config.json`
6. Deploy by running `cf push`