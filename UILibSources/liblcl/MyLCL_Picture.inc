
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------

function Picture_Create: TPicture; extdecl;
begin
  Result :=  TPicture.Create;
end;

procedure Picture_Free(AObj: TPicture); extdecl;
begin
  AObj.Free;
end;

procedure Picture_LoadFromFile(AObj: TPicture; Filename: PChar); extdecl;
begin
  AObj.LoadFromFile(Filename);
end;

procedure Picture_SaveToFile(AObj: TPicture; Filename: PChar); extdecl;
begin
  AObj.SaveToFile(Filename);
end;

procedure Picture_LoadFromStream(AObj: TPicture; Stream: TStream); extdecl;
begin
  AObj.LoadFromStream(Stream);
end;

procedure Picture_SaveToStream(AObj: TPicture; Stream: TStream); extdecl;
begin
  AObj.SaveToStream(Stream);
end;

procedure Picture_Assign(AObj: TPicture; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function Picture_GetNamePath(AObj: TPicture): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Picture_ClassType(AObj: TPicture): TClass; extdecl;
begin
  Result :=  AObj.ClassType;
end;

function Picture_ClassName(AObj: TPicture): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Picture_InstanceSize(AObj: TPicture): Integer; extdecl;
begin
  Result :=  AObj.InstanceSize;
end;

function Picture_InheritsFrom(AObj: TPicture; AClass: TClass): LongBool; extdecl;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function Picture_Equals(AObj: TPicture; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Picture_GetHashCode(AObj: TPicture): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Picture_ToString(AObj: TPicture): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Picture_GetBitmap(AObj: TPicture): TBitmap; extdecl;
begin
  Result :=  AObj.Bitmap;
end;

procedure Picture_SetBitmap(AObj: TPicture; AValue: TBitmap); extdecl;
begin
  AObj.Bitmap := AValue;
end;

function Picture_GetGraphic(AObj: TPicture): TGraphic; extdecl;
begin
  Result :=  AObj.Graphic;
end;

procedure Picture_SetGraphic(AObj: TPicture; AValue: TGraphic); extdecl;
begin
  AObj.Graphic := AValue;
end;

function Picture_GetHeight(AObj: TPicture): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

function Picture_GetIcon(AObj: TPicture): TIcon; extdecl;
begin
  Result :=  AObj.Icon;
end;

procedure Picture_SetIcon(AObj: TPicture; AValue: TIcon); extdecl;
begin
  AObj.Icon := AValue;
end;

function Picture_GetWidth(AObj: TPicture): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure Picture_SetOnChange(AObj: TPicture; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnChange}
  {$DEFINE EventMethodName := OnChange}
  EventMethodCode
end;

function Picture_StaticClassType: TClass; extdecl;
begin
  Result := TPicture;
end;

exports
  Picture_Create,
  Picture_Free,
  Picture_LoadFromFile,
  Picture_SaveToFile,
  Picture_LoadFromStream,
  Picture_SaveToStream,
  Picture_Assign,
  Picture_GetNamePath,
  Picture_ClassType,
  Picture_ClassName,
  Picture_InstanceSize,
  Picture_InheritsFrom,
  Picture_Equals,
  Picture_GetHashCode,
  Picture_ToString,
  Picture_GetBitmap,
  Picture_SetBitmap,
  Picture_GetGraphic,
  Picture_SetGraphic,
  Picture_GetHeight,
  Picture_GetIcon,
  Picture_SetIcon,
  Picture_GetWidth,
  Picture_SetOnChange,
  Picture_StaticClassType;

