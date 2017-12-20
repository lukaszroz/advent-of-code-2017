p = day20(:,1:3);
v = day20(:,4:6);
a = day20(:,7:9);

#scatter3(p(:,1),p(:,2),p(:,3));
#axis equal;
#aa = axis;
#axis(aa);

c = [zeros(size(p, 1), 2) ones(size(p, 1), 1)];
data = {p c};

for ii = 1:100
    v += a;
    pn = p + v;
    for j = 1:9
      pp = p + j * (pn - p) / 10;      
      data((ii - 1) * 10 + j,:) =  {pp c};
    end
    p = pn;
    data(ii*10,:) =  {p c};
    c_rows = zeros(size(p,1), 1);
    for j = 1:size(p,1)
     count = sum((p(j, :) == p), 2);
      c_rows += count == 3;
    end
    
    c_rows = c_rows > 1;
    
    v(c_rows, :) = [0];
    a(c_rows, :) = [0];
    c(c_rows, 3) = 0;
    c(c_rows, 1) = 1;
 end