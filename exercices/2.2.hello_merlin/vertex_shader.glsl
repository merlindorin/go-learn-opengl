#version 330 core

layout (location = 0) in vec3 position;
layout (location = 1) in vec3 color;

out vec3 ourColor; // Output a color to the fragment shader

void main()
{
    gl_Position = vec4(position, 1.0);
    ourColor = vec3(color.x, color.y, .5);
}