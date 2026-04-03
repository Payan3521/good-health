using Microsoft.AspNetCore.Mvc;
namespace notification_microservice.Controllers;

[ApiController]
[Route("/")]
public class TestController : ControllerBase
{
    [HttpGet]
    public IActionResult Get()
    {
        return Ok("Notification Microservice is running with C# and .NET");
    }
}