-- Inserting questions
INSERT INTO "question" ("text", "teacher_id") VALUES
    ('Allergiya nima?', '1'),
    ('Biologik allergenlarga kiradi', '1'),
    ('Gumoral shok sababi:', '1'),
    ('Gemotransfuziyadan keyin tekshirish shart emas:', '1'),
    ('Normada chaqaloqlarning nafas olish soni 1 daqiqada qancha?', '1'),
    ('Bronxial astmaga xarakterli bo''lgan simtomlar:', '1'),
    ('Termik kuyishlarda asosiy shoshilinch tibbiy yordam:', '1'),
    ('Yurak - o''pka reanimatsiyasining asoratlari:', '1'),
    ('Miokard infarktini medikamentoz davolashda eng avval qo''llaniladigan preparatlarni ko’rsating:', '1'),
    ('Gipertoniya kasalligi bilan kasallangan bemorda obyektiv tekshiruvda quyidagilar aniqlandi:', '1'),
    ('Oshqozon yarasi perforasiyasiga xos simptomlar:', '1'),
    ('Tungi siydik tutolmaslik nima deb ataladi?', '1'),
    ('Vaqtinchalik qon to''xtatish usuli', '1'),
    ('Yurak - o''pka reanimasiyasining o''tkazilmaydigan holatlar', '1'),
    ('Qon ketishini to''xtatish usullari farqlanadi:', '1'),
    ('Qon ketishining og''ir a’soratini aniqlang', '1'),
    ('Qaysi kasalliklarda qon ivishi sekinlashadi:', '1'),
    ('O''pkadan qon ketish belgilari:', '1'),
    ('O''tkir siydik tutilishida shoshilinch yordam:', '1'),
    ('Gipertonik krizda kaysi dori qo''llanilmaydi?', '1');

-- Inserting choices
-- For each question, only the first choice is marked as correct (is_correct = true)
INSERT INTO "choices" ("question_id", "choice_text", "is_correct") VALUES
    ('1', '1allergenga nisbatan organizm sezgirligi ortishi natijasidagi reaksiya', true),
    ('1', 'infeksion allergik kasallik', false),
    ('1', 'ortirilgan immunitet tanqisligi', false),
    ('1', 'surunkali kasallik asorati', false),

    ('2', '1mikrob, virus, zamburug‘, gelmint', true),
    ('2', 'antibiotik, vitamin, sulfanilamid', false),
    ('2', 'uydan, gilam, to‘shakdan chiqqan chang', false),
    ('2', 'sut, tuxum, pomidor, qulupnay, shokolad', false),

    ('3', '1xos bo''lmagan guruh qonni quyish', true),
    ('3', 'zararli ekologik omillar', false),
    ('3', 'ichki organlar kasalligi', false),
    ('3', 'nurlanish kirishidan', false),

    ('4', '1Antropometriya', true),
    ('4', 'termometriya', false),
    ('4', 'tonometriya', false),
    ('4', 'siydik analizi', false),

    ('5', '25-50', true),
    ('5', '15-30', false),
    ('5', '12-20', false),
    ('5', '8-10', false),

    ('6', '1nafas chiqarishini qiyinlashishi bilan kechuvchi bo’g''ilish', true),
    ('6', 'nafas olishini qiyinlashishi bilan kechuvchi bug''ilish', false),
    ('6', 'nafas olish yuzaki va tez-tez, burun-lab uchburchagi ko''karishi', false),
    ('6', 'nafas olishini qiyinlashishi, o''pkada nam mayda pufakli xirillashlar', false),

    ('7', '1og''riqsizlantirish', true),
    ('7', 'tinchlantirish', false),
    ('7', 'gipotenziv terapiya', false),
    ('7', 'gipertenziv terapiya', false),

    ('8', '1qovurg''alarning sinishi; ko''krak qafasida qon to''planishi; singan qovurg''a o''pka, yurakni jaroxatlantirishi; jigarni ezilishi; oshqozondan qusuq massalarini qaytib chiqishi', true),
    ('8', 'tashqi uyqu arteriyasida mexanik puls to''lqinining paydo bo''lishi; teri rangi qizarishi; qo''z qorachiqlarni qisqarishi', false),
    ('8', 'qovurg''alarning sinishi; ko''krak qafasida qon to''plinishi; singan qovurg''a o''pka, yurakni jaroxatlantirishi; buyraqlarni ezilishi; oshqozon qusuk massalarini qaytib chiqishi', false),
    ('8', 'qovurg''alarning sinishi; ko''krak qafasida qon to''planishi; singan qovurg''a o''pka, yurakni jaroxatlantirishi; oshqozon qusuq massalarini qaytib chiqishi', false),

    ('9', '1nitratlar, antikoagulyantlar, B–adrenoblokatorlar', true),
    ('9', 'antigistamin preparatlar', false),
    ('9', 'a-adrenoblokatorlar', false),
    ('9', 'strofantin, eufillin, panangin', false),

    ('10', '1pulsni taranglashishi, AQB oshishi, yurakning kattalashishi', true),
    ('10', 'pulsni kuchsizligi, lablar sianozi, yurak auskultatsiyada yurak chuqqisida sistolik shovqin', false),
    ('10', 'aritmik puls, oyoqlarda shish, AQBning ko’tarilishi', false),
    ('10', 'yuz va tanada shish, teri rangi oqargan, AQB ko’tarilgan', false),

    ('11', '1xanjarsimon og''riq epigastral soxada', true),
    ('11', 'safro bilan kayd qilish', false),
    ('11', 'axlatni oqish rangda bo''lishi', false),
    ('11', 'terini qichishi', false),

    ('12', '1Enurez', true),
    ('12', 'anuriya', false),
    ('12', 'oliguriya', false),
    ('12', 'ishuriya', false),

    ('13', '1jgut qo''yish', true),
    ('13', 'termik', false),
    ('13', 'mexanik', false),
    ('13', 'ximiyaviy', false),

    ('14', '1maxsus jiton bo''lganlar; ko''krak qafasi singanda; diniy tarabdan ruxsat bermaganda; hayotiy nomunosib jaroxatlarda', true),
    ('14', 'qovurg''alarning sinishida; ko''krak qafasida qon to’planishida; jigarni ezilishida', false),
    ('14', 'qovurg''alarning sinishida; ko''krak qafasida qon to’planishida; singan qovurg''a o''pka, yurakni jaroxatlantirishida', false),
    ('14', 'ko''krak qafasi singanda; singan qovurg''a o''pka, yurakni jaroxatlantirishida; jigarni ezilishida;', false),

    ('15', '1vaqtincha va uzil-kesil', true),
    ('15', 'fizik, kimyoviy', false),
    ('15', 'mexanik, biologik', false),
    ('15', 'ishonchli, ishonchsiz', false),

    ('16', '1gemorragik shok', true),
    ('16', 'xushdan ketish', false),
    ('16', 'emboliya', false),
    ('16', 'pnevmotoraks', false),

    ('17', '1gemofiliya, nur kasalligi, geparinni ko''p dozada yuborish', true),
    ('17', 'Verlgof kasalligi, oshqozon yara kasalligi, ateroskleroz', false),
    ('17', 'Mellori-Veys sindromi, geparinni ko''p dozada yuborishi', false),
    ('17', 'vitamin K-ning normadan oshishi, nur kasalligi, gemofiliya', false),

    ('18', '1qonning qizil rangda ko''piksimon, yo''tal bilan bog''lik bo''lishi', true),
    ('18', 'eritrositlar va AQBni tushib ketishi, melena', false),
    ('18', '“kofe rangi"dagi qusish, yo''tal bilan bog''lik bo''lishi', false),
    ('18', '“kofe rangi"dagi qusish, qonning qizil rangda bo''lishi', false),

    ('19', 'kateter qo’yish', true),
    ('19', 'og''riqsizlantirish', false),
    ('19', 'issiq grelka quyish', false),
    ('19', 'spazmolitiklarni qo''llash', false),

    ('20', '1Dofamin', false),
    ('20', 'magneziya sulfat', true),
    ('20', 'dibazol', false),
    ('20', 'klofelin', false);
