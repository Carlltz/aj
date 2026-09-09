{
  lib,
  buildGoModule,
}:

buildGoModule {
  pname = "aj";
  version = "0.1.0";

  src = lib.cleanSource ./.;

  vendorHash = "sha256-EZCSmKxsSwBMJHR6I3MvV2cWjc43z4fJ8fjCuDzylsU=";

  ldflags = [
    "-s"
    "-w"
  ];

  meta = {
    description = "Shell tool that automatically fixes mistakes in your last command";
    homepage = "https://github.com/Carlltz/aj";
    license = lib.licenses.mit;
    mainProgram = "aj";
    platforms = lib.platforms.unix;
  };
}
