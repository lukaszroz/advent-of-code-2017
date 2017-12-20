figure(1,"visible","off");
P = cell2mat(data(1, 1));
C = cell2mat(data(1, 2));
scatter3(P(:,1),P(:,2),P(:,3),[],C);
axis equal;
aa = axis;
axis(aa);
angle = -37.5;
view(angle, 30)
drawnow;
frame = 1;
filename = strcat(num2str(frame),'.png');
print(filename, '-dpng');
frame +=1
zoom = 1;

for i = 1:500
  hold off
  P = cell2mat(data(i, 1));
  C = cell2mat(data(i, 2));
  if i < 600
    angle = -37.5 + i * 30 / 500;
  else
    angle = -37.5 + 35 - (i - 500) * 30 / 500;
  end
  
  #zoom in
  if i == 100
    for j = 1:60
      zoom = -0.016101694915254	* j + 1.0161016949153;
      scatter3(P(:,1),P(:,2),P(:,3),[],C);
      axis equal;
      axis(aa * zoom);
      view(angle, 30)
      drawnow;
      filename = strcat(num2str(frame),'.png');
      print(filename, '-dpng');
      frame +=1
    end
  end
  
  if i == 350
    for j = 1:60
      zoom = 0.033050847457627 * j + 0.016949152542373;
      scatter3(P(:,1),P(:,2),P(:,3),[],C);
      axis equal;
      axis(aa * zoom);
      view(angle, 30)
      drawnow;      
      filename = strcat(num2str(frame),'.png');
      print(filename, '-dpng');
      frame +=1
    end
  end
  

  scatter3(P(:,1),P(:,2),P(:,3),[],C);
  axis equal;
  axis(aa * zoom);
  view(angle, 30)
  drawnow;
  filename = strcat(num2str(frame),'.png');
  print(filename, '-dpng');
  frame +=1
end