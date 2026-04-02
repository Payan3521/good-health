var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

// Tu endpoint de "Hola Mundo" (equivalente a @app.route('/'))
app.MapGet("/", () => "¡Hola Mundo desde .NET!");

// Un segundo endpoint tipo JSON
app.MapGet("/api/status", () => new { mensaje = "Microservicio activo", version = 1.0 });

app.Run();
