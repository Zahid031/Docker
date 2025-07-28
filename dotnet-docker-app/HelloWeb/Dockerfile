FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
WORKDIR /app
COPY . .
RUN dotnet publish -c Release -o /app/publish

FROM mcr.microsoft.com/dotnet/aspnet:8.0 AS final
WORKDIR /app
COPY --from=build /app/publish .
ENTRYPOINT ["dotnet", "HelloWeb.dll"]
EXPOSE 8080






# FROM mcr.microsoft.com/dotnet/sdk:8.0 AS base
# WORKDIR /app


# FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
# WORKDIR /app
# COPY . .
# RUN dotnet publish -c Release -o /app/publish

# FROM build AS final
# WORKDIR /app
# COPY --from=build /app/publish .
# ENTRYPOINT ["dotnet", "HelloWeb.dll"]
# EXPOSE 8080

